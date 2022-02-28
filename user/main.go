package main

import (
	"github.com/Reoneks/microservice_chat/user/config"
	"github.com/Reoneks/microservice_chat/user/user"

	"github.com/Reoneks/microservice_chat/proto"

	"github.com/asim/go-micro/v3"
)

func main() {
	cfg := config.GetConfig()
	db, err := config.NewDB(cfg.MongoUrl)
	if err != nil {
		panic(err)
	}

	userRep := user.NewUserRepository(db, cfg.DBName, cfg.Collection)
	userService := user.NewUserService(userRep)
	microService := micro.NewService(micro.Name(cfg.ServiceName), micro.Address(cfg.MicroServiceAddress))
	microService.Init()
	if err := proto.RegisterUserHandler(microService.Server(), &user.UserMicro{UserService: userService}); err != nil {
		panic(err)
	}
	if err := microService.Run(); err != nil {
		panic(err)
	}
}
