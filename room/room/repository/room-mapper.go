package repository

import (
	room "chatex/room/room/room_interface"
)

func FromRoomsDto(roomDto room.RoomsDto) room.Rooms {
	return room.Rooms(roomDto)
}

func FromRoomsDtos(roomsDtos []room.RoomsDto) (rooms []room.Rooms) {
	for _, dto := range roomsDtos {
		rooms = append(rooms, FromRoomsDto(dto))
	}
	return
}

func ToRoomsDto(roomToConvert room.Rooms) room.RoomsDto {
	return room.RoomsDto(roomToConvert)
}

func ToRoomsDtos(rooms []room.Rooms) (RoomsDtos []room.RoomsDto) {
	for _, dto := range rooms {
		RoomsDtos = append(RoomsDtos, ToRoomsDto(dto))
	}
	return
}
