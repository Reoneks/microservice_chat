package main

import (
	"net/http"

	"github.com/Reoneks/microservice_chat/api-gateway/clients"
	"github.com/Reoneks/microservice_chat/api-gateway/config"
	"github.com/Reoneks/microservice_chat/api-gateway/connector"
	"github.com/Reoneks/microservice_chat/api-gateway/server"
	"github.com/gorilla/websocket"

	"github.com/Reoneks/microservice_chat/proto"

	"github.com/asim/go-micro/v3"
)

const (
	readBufferSize  = 1024
	writeBufferSize = 1024
)

func main() {
	cfg := config.GetConfig()
	log, err := config.NewLogger(cfg.LogLevel)
	if err != nil {
		panic(err)
	}

	var upgrader = &websocket.Upgrader{
		ReadBufferSize:  readBufferSize,
		WriteBufferSize: writeBufferSize,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	authService := micro.NewService(micro.Address(cfg.AuthServiceADDR))
	authService.Init()
	userService := micro.NewService(micro.Address(cfg.UserServiceADDR))
	userService.Init()
	roomService := micro.NewService(micro.Address(cfg.RoomServiceADDR))
	roomService.Init()
	messagesService := micro.NewService(micro.Address(cfg.MessageServiceADDR))
	messagesService.Init()

	auth := proto.NewAuthService(cfg.AuthServiceName, authService.Client())
	user := proto.NewUserService(cfg.UserServiceName, userService.Client())
	room := proto.NewRoomsService(cfg.RoomServiceName, roomService.Client())
	messages := proto.NewMessagesService(cfg.MessageServiceName, messagesService.Client())

	authMicroservice := clients.NewAuthMicroservice(auth)
	userMicroservice := clients.NewUserMicroservice(user, auth)
	roomMicroservice := clients.NewRoomMicroservice(room, cfg.DefaltLimit, cfg.DefaltOffset)
	messagesMicroservice := clients.NewMessagesMicroservice(messages, cfg.DefaltLimit, cfg.DefaltOffset)

	amqp := config.GetAMQP(&cfg, nil, nil, nil)
	ch, err := config.StartRabbitMQ(cfg.RabbitMQUrl)
	if err != nil {
		panic(err)
	}

	_, err = amqp.GetSendQueue(ch)
	if err != nil {
		panic(err)
	}

	msgs, err := amqp.GetReceiveChan(ch)
	if err != nil {
		panic(err)
	}

	connect := connector.NewWSConnector(log, ch, &cfg, msgs, room)

	server := server.NewHTTPServer(
		log,
		cfg.Addr,
		authMicroservice,
		userMicroservice,
		roomMicroservice,
		messagesMicroservice,
		auth,
		cfg.ApiGatewaySubscribeName,
		connect,
		upgrader,
	)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
