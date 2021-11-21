package main

import (
	"github.com/Reoneks/microservice_chat/auth/auth"
	"github.com/Reoneks/microservice_chat/auth/config"

	"github.com/Reoneks/microservice_chat/proto"

	promwrapper "github.com/asim/go-micro/plugins/wrapper/monitoring/prometheus/v3"
	"github.com/asim/go-micro/v3"
)

func main() {
	cfg := config.GetConfig()

	db, err := config.NewDB(cfg.DSN, cfg.MigrationURL)
	if err != nil {
		panic(err)
	}

	jwt := cfg.NewJWT()

	authRepository := auth.NewAuthRepository(db)
	authService := auth.NewAuthService(authRepository, jwt)

	service := micro.NewService(
		micro.Name(cfg.ServiceName),
		micro.WrapHandler(promwrapper.NewHandlerWrapper(
			promwrapper.ServiceName(cfg.ServiceName),
			promwrapper.ServiceID("auth"),
		)),
	)
	service.Init()
	err = proto.RegisterAuthServiceHandler(
		service.Server(),
		auth.NewAuth(
			authService,
			proto.NewUserService(cfg.UserServiceName, service.Client()),
			jwt,
		),
	)
	if err != nil {
		panic(err)
	}
	if err := service.Run(); err != nil {
		panic(err)
	}
}
