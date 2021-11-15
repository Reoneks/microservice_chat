package service

import (
	"chatex/proto"
	room "chatex/room/room/room_interface"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromRoomsStruct(roomsStruct *proto.RoomStruct) *room.Rooms {
	return &room.Rooms{
		ID:        roomsStruct.ID,
		Name:      roomsStruct.Name,
		Status:    roomsStruct.Status,
		CreatedBy: roomsStruct.CreatedBy,
		CreatedAt: roomsStruct.CreatedAt.AsTime(),
	}
}

func ToRoomsStructRSP(room *room.Rooms, rsp *proto.RoomStructResponse) {
	rsp.Room.ID = room.ID
	rsp.Room.Name = room.Name
	rsp.Room.Status = room.Status
	rsp.Room.CreatedBy = room.CreatedBy
	rsp.Room.CreatedAt = timestamppb.New(room.CreatedAt)

	rsp.Status.Ok = true
	rsp.Status.Error = ""
}

func ToRoomsStructsRSP(rooms []room.Rooms, rsp *proto.GetRoomsResponse) {
	for _, dto := range rooms {
		var data *proto.RoomStructResponse
		ToRoomsStructRSP(&dto, data)
		rsp.Rooms = append(rsp.Rooms, data.Room)
	}

	rsp.Status.Ok = true
	rsp.Status.Error = ""
}

func FromFilter(filter *proto.Filter) *room.RoomsFilter {
	return &room.RoomsFilter{
		IDs: filter.RoomIDs,
	}
}
