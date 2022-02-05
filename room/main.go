package main

import (
	"github.com/Reoneks/microservice_chat/proto"
	"github.com/Reoneks/microservice_chat/room/config"
	"github.com/Reoneks/microservice_chat/room/room"

	"github.com/asim/go-micro/v3"
)

func main() {
	cfg := config.GetConfig()
	db, err := config.NewDB(cfg.MongoUrl)
	if err != nil {
		panic(err)
	}

	roomRep := room.NewRoomRepository(db, cfg.DBName, cfg.RoomCollection)
	roomUserRep := room.NewRoomUsersRepository(db, cfg.DBName, cfg.RoomUserCollection, cfg.RoomCollection)

	roomService := room.NewRoomService(roomRep, roomUserRep)
	microService := micro.NewService(micro.Name(cfg.ServiceName))
	microService.Init()

	if err := proto.RegisterRoomsHandler(microService.Server(), &room.RoomsMicro{RoomService: roomService}); err != nil {
		panic(err)
	}

	if err := microService.Run(); err != nil {
		panic(err)
	}
}
