package room

import (
	"context"
	"encoding/json"

	"github.com/Reoneks/microservice_chat/proto"
	"github.com/Reoneks/microservice_chat/room/model"
)

type RoomsMicro struct {
	RoomService IRoomService
}

func (u *RoomsMicro) GetRoom(ctx context.Context, req *proto.RoomID, rsp *proto.RoomStructResponse) error {
	room, err := u.RoomService.GetRoom(req.ID)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	bytes, err := json.Marshal(room)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	rsp.Room.RoomInfo = bytes
	rsp.Status.Ok = true
	return nil
}

func (u *RoomsMicro) GetRooms(ctx context.Context, req *proto.Filter, rsp *proto.GetRoomsResponse) error {
	rooms, err := u.RoomService.GetRooms(&RoomsFilter{IDs: req.RoomIDs})
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	var resp []*proto.RoomStruct

	for _, room := range rooms {
		bytes, err := json.Marshal(room)
		if err != nil {
			rsp.Status.Ok = false
			rsp.Status.Error = err.Error()
			return err
		}

		resp = append(resp, &proto.RoomStruct{
			RoomInfo: bytes,
		})
	}

	rsp.Rooms = resp
	rsp.Status.Ok = true
	return nil
}

func (u *RoomsMicro) CreateRoom(ctx context.Context, req *proto.RoomStruct, rsp *proto.RoomStructResponse) error {
	var data model.RoomsDto
	err := json.Unmarshal(req.RoomInfo, &data)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	room, err := u.RoomService.CreateRoom(&data)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	bytes, err := json.Marshal(room)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	rsp.Room.RoomInfo = bytes
	rsp.Status.Ok = true
	return nil
}

func (u *RoomsMicro) DeleteRoom(ctx context.Context, req *proto.DeleteRequest, rsp *proto.Status) error {
	err := u.RoomService.DeleteRoom(req.RoomID, req.UserID)
	if err != nil {
		rsp.Ok = false
		rsp.Error = err.Error()
		return err
	}

	rsp.Ok = true
	return nil
}

func (u *RoomsMicro) UpdateRoom(ctx context.Context, req *proto.UpdateRequest, rsp *proto.RoomStructResponse) error {
	var data model.RoomsDto
	err := json.Unmarshal(req.Room.RoomInfo, &data)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	room, err := u.RoomService.UpdateRoom(&data, req.UserID)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	bytes, err := json.Marshal(room)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	rsp.Room.RoomInfo = bytes
	rsp.Status.Ok = true
	return nil
}

func (u *RoomsMicro) AddUsers(ctx context.Context, req *proto.AddUsersRequest, rsp *proto.Status) error {
	err := u.RoomService.AddUsers(req.RoomID, req.UserID, req.UserIDs)
	if err != nil {
		rsp.Ok = false
		rsp.Error = err.Error()
		return err
	}

	rsp.Ok = true
	return nil
}
