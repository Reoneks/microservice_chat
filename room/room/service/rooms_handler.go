package service

import (
	"chatex/proto"
	room "chatex/room/room/room_interface"
	"context"
)

type RoomsMicro struct {
	RoomService room.RoomService
}

func (u *RoomsMicro) GetRoom(ctx context.Context, req *proto.RoomID, rsp *proto.RoomStructResponse) error {
	room, err := u.RoomService.GetRoom(req.ID)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	ToRoomsStructRSP(room, rsp)
	return nil
}

func (u *RoomsMicro) GetRooms(ctx context.Context, req *proto.Filter, rsp *proto.GetRoomsResponse) error {
	rooms, err := u.RoomService.GetRooms(FromFilter(req))
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	ToRoomsStructsRSP(rooms, rsp)
	return nil
}

func (u *RoomsMicro) CreateRoom(ctx context.Context, req *proto.RoomStruct, rsp *proto.RoomStructResponse) error {
	room, err := u.RoomService.CreateRoom(FromRoomsStruct(req))
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	ToRoomsStructRSP(room, rsp)
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
	room, err := u.RoomService.UpdateRoom(FromRoomsStruct(req.Room), req.UserID)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	ToRoomsStructRSP(room, rsp)
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
