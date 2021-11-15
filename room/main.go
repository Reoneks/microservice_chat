package main

import (
	"chatex/proto"
	"chatex/room/config"
	"chatex/room/room/repository"
	"chatex/room/room/service"
	userRep "chatex/room/room_user/repository"

	"github.com/asim/go-micro/v3"
)

func main() {
	cfg := config.GetConfig()
	db, err := config.NewDB(cfg.DSN, cfg.MigrationURL)
	if err != nil {
		panic(err)
	}

	roomRep := repository.NewRoomRepository(db)
	roomUserRep := userRep.NewRoomUsersRepository(db)

	roomService := service.NewRoomService(roomRep, roomUserRep)
	microService := micro.NewService(micro.Name(cfg.ServiceName))
	microService.Init()

	//TODO: add room handler
	if err := proto.RegisterRoomsHandler(microService.Server(), &service.RoomsMicro{RoomService: roomService}); err != nil {
		panic(err)
	}

	if err := microService.Run(); err != nil {
		panic(err)
	}
}
