package repository

import (
	room "chatex/room/room/room_interface"
	room_user "chatex/room/room_user/room_user_interface"

	gm "gorm.io/gorm"
)

type RoomUsersRepository interface {
	GetRoomsByUserId(id string) ([]room.RoomsDto, error)
	CreateUserRoomConnection(user room_user.RoomUsersDto) (*room_user.RoomUsersDto, error)
}

type RoomUsersRepositoryImpl struct {
	db *gm.DB
}

func (r *RoomUsersRepositoryImpl) GetRoomsByUserId(id string) (rooms []room.RoomsDto, err error) {
	var roomUsers []room_user.RoomUsersDto
	if err = r.db.Where("user_id = ?", id).Find(&roomUsers).Error; err != nil {
		return
	}
	for _, user := range roomUsers {
		var roomStruct *room.RoomsDto
		if err = r.db.Where("id = ?", user.RoomID).First(&roomStruct).Error; err != nil {
			rooms = nil
			return
		}
		rooms = append(rooms, *roomStruct)
	}
	return
}

func (r *RoomUsersRepositoryImpl) CreateUserRoomConnection(user room_user.RoomUsersDto) (*room_user.RoomUsersDto, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewRoomUsersRepository(db *gm.DB) RoomUsersRepository {
	return &RoomUsersRepositoryImpl{
		db,
	}
}
