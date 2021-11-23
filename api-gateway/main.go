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
	url := cfg.ServerAddress()
	log, err := config.NewLogger(cfg.LogLevel)
	if err != nil {
		panic(err)
	}

	var upgrader = &websocket.Upgrader{
		ReadBufferSize:  readBufferSize,
		WriteBufferSize: writeBufferSize,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	service := micro.NewService()
	service.Init()

	auth := proto.NewAuthService(cfg.AuthServiceName, service.Client())
	user := proto.NewUserService(cfg.UserServiceName, service.Client())
	room := proto.NewRoomsService(cfg.RoomServiceName, service.Client())
	messages := proto.NewMessagesService(cfg.MessageServiceName, service.Client())

	authMicroservice := clients.NewAuthMicroservice(auth)
	userMicroservice := clients.NewUserMicroservice(user, auth)
	roomMicroservice := clients.NewRoomMicroservice(room)
	messagesMicroservice := clients.NewMessagesMicroservice(messages)

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

	connect := connector.NewWSConnector(log, ch, &cfg, msgs)

	server := server.NewHTTPServer(
		log,
		url,
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
