package main

import (
	"github.com/Reoneks/microservice_chat/api-gateway/clients"
	"github.com/Reoneks/microservice_chat/api-gateway/config"
	"github.com/Reoneks/microservice_chat/api-gateway/server"

	"github.com/Reoneks/microservice_chat/proto"

	"github.com/asim/go-micro/v3"
)

func main() {
	cfg := config.GetConfig()
	url := cfg.ServerAddress()
	log, err := config.NewLogger(cfg.LogLevel)
	if err != nil {
		panic(err)
	}

	service := micro.NewService()
	service.Init()

	auth := proto.NewAuthService(cfg.AuthServiceName, service.Client())
	user := proto.NewUserService(cfg.UserServiceName, service.Client())
	room := proto.NewRoomsService(cfg.RoomServiceName, service.Client())

	authMicroservice := clients.NewAuthMicroservice(auth)
	userMicroservice := clients.NewUserMicroservice(user, auth)
	roomMicroservice := clients.NewRoomMicroservice(room)

	server := server.NewHTTPServer(
		log,
		url,
		authMicroservice,
		userMicroservice,
		roomMicroservice,
		auth,
		cfg.ApiGatewaySubscribeName,
	)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
