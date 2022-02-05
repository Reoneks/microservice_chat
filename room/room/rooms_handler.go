package room

import (
	"context"
	"encoding/json"

	"github.com/Reoneks/microservice_chat/proto"
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

	rsp.Room = bytes
	rsp.Status.Ok = true
	return nil
}

func (u *RoomsMicro) CreateRoom(ctx context.Context, req *proto.RoomStruct, rsp *proto.RoomStructResponse) error {
	data := make(map[string]interface{})
	err := json.Unmarshal(req.RoomInfo, &data)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	room, err := u.RoomService.CreateRoom(data)
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

	rsp.Room = bytes
	rsp.Status.Ok = true
	return nil
}

func (u *RoomsMicro) DeleteRoom(ctx context.Context, req *proto.DeleteRequest, rsp *proto.Status) error {
	err := u.RoomService.DeleteRoom(req.RoomID)
	if err != nil {
		rsp.Ok = false
		rsp.Error = err.Error()
		return err
	}

	rsp.Ok = true
	return nil
}

func (u *RoomsMicro) UpdateRoom(ctx context.Context, req *proto.UpdateRequest, rsp *proto.RoomStructResponse) error {
	data := make(map[string]interface{})
	err := json.Unmarshal(req.Room, &data)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	room, err := u.RoomService.UpdateRoom(data)
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

	rsp.Room = bytes
	rsp.Status.Ok = true
	return nil
}

func (u *RoomsMicro) AddUsers(ctx context.Context, req *proto.AddUsersRequest, rsp *proto.Status) error {
	err := u.RoomService.AddUsers(req.RoomID, req.UserIDs)
	if err != nil {
		rsp.Ok = false
		rsp.Error = err.Error()
		return err
	}

	rsp.Ok = true
	return nil
}

func (u *RoomsMicro) DeleteUsers(ctx context.Context, req *proto.AddUsersRequest, rsp *proto.Status) error {
	err := u.RoomService.DeleteUsers(req.RoomID, req.UserIDs)
	if err != nil {
		rsp.Ok = false
		rsp.Error = err.Error()
		return err
	}

	rsp.Ok = true
	return nil
}

func (u *RoomsMicro) GetAllRooms(ctx context.Context, req *proto.GetAllRoomsRequest, rsp *proto.RoomStructResponse) error {
	rooms, err := u.RoomService.GetAllRooms(req.UserID, req.Limit, req.Offset)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	bytes, err := json.Marshal(rooms)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	rsp.Room = bytes
	rsp.Status.Ok = true
	return nil
}
