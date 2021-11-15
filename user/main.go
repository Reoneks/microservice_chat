package main

import (
	"user_service/config"
	"user_service/user"

	"github.com/Reoneks/microservice_chat/proto"

	"github.com/asim/go-micro/v3"
)

func main() {
	cfg := config.GetConfig()
	db, err := config.NewDB(cfg.DSN, cfg.MigrationURL)
	if err != nil {
		panic(err)
	}

	userRep := user.NewUserRepository(db)
	userService := user.NewUserService(userRep)
	microService := micro.NewService(micro.Name(cfg.ServiceName))
	microService.Init()
	if err := proto.RegisterUserHandler(microService.Server(), &user.UserMicro{UserService: userService}); err != nil {
		panic(err)
	}
	if err := microService.Run(); err != nil {
		panic(err)
	}
}
