package main

import (
	"github.com/Reoneks/microservice_chat/auth/auth"
	"github.com/Reoneks/microservice_chat/auth/config"

	"github.com/Reoneks/microservice_chat/proto"

	"github.com/asim/go-micro/v3"
)

func main() {
	cfg := config.GetConfig()

	db, err := config.NewDB(cfg.MongoUrl)
	if err != nil {
		panic(err)
	}

	jwt := cfg.NewJWT()

	authRepository := auth.NewAuthRepository(db, cfg.DBName, cfg.Collection)
	authService := auth.NewAuthService(authRepository, jwt)

	userService := micro.NewService(micro.Address(cfg.UserServiceADDR))
	userService.Init()
	authMicroService := micro.NewService(micro.Name(cfg.ServiceName), micro.Address(cfg.MicroServiceAddress))
	authMicroService.Init()

	err = proto.RegisterAuthServiceHandler(
		authMicroService.Server(),
		auth.NewAuth(
			authService,
			proto.NewUserService(cfg.UserServiceName, userService.Client()),
			jwt,
		),
	)
	if err != nil {
		panic(err)
	}
	if err := authMicroService.Run(); err != nil {
		panic(err)
	}
}
