package main

import (
	"proto"
	"user/config"
	"user_service/user/repository"
	"user_service/user/service"

	"github.com/asim/go-micro/v3"
)

func main() {
	cfg := config.GetConfig()
	db, err := config.NewDB(cfg.DSN, cfg.MigrationURL)
	if err != nil {
		panic(err)
	}

	userRep := repository.NewUserRepository(db)
	userService := service.NewUserService(userRep)
	microService := micro.NewService(micro.Name(cfg.ServiceName))
	microService.Init()
	if err := proto.RegisterUserHandler(microService.Server(), &service.UserMicro{UserService: userService}); err != nil {
		panic(err)
	}
	if err := microService.Run(); err != nil {
		panic(err)
	}
}
