package main

import (
	"github.com/Reoneks/microservice_chat/proto"
	"github.com/Reoneks/microservice_chat/room/config"
	"github.com/Reoneks/microservice_chat/room/room"
	userRep "github.com/Reoneks/microservice_chat/room/room_user"

	"github.com/asim/go-micro/v3"
)

func main() {
	cfg := config.GetConfig()
	db, err := config.NewDB(cfg.DSN, cfg.MigrationURL)
	if err != nil {
		panic(err)
	}

	roomRep := room.NewRoomRepository(db)
	roomUserRep := userRep.NewRoomUsersRepository(db)

	roomService := room.NewRoomService(roomRep, roomUserRep)
	microService := micro.NewService(micro.Name(cfg.ServiceName))
	microService.Init()

	//TODO: add room handler
	if err := proto.RegisterRoomsHandler(microService.Server(), &room.RoomsMicro{RoomService: roomService}); err != nil {
		panic(err)
	}

	if err := microService.Run(); err != nil {
		panic(err)
	}
}
