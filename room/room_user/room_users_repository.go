package room_user

import (
	"github.com/Reoneks/microservice_chat/room/model"
	gm "gorm.io/gorm"
)

type RoomUsersRepository interface {
	GetRoomsByUserId(id string) ([]model.RoomsDto, error)
	CreateUserRoomConnection(user model.RoomUsersDto) (*model.RoomUsersDto, error)
}

type RoomUsersRepositoryImpl struct {
	db *gm.DB
}

func (r *RoomUsersRepositoryImpl) GetRoomsByUserId(id string) (rooms []model.RoomsDto, err error) {
	var roomUsers []model.RoomUsersDto
	if err = r.db.Where("user_id = ?", id).Find(&roomUsers).Error; err != nil {
		return
	}
	for _, user := range roomUsers {
		var roomStruct *model.RoomsDto
		if err = r.db.Where("id = ?", user.RoomID).First(&roomStruct).Error; err != nil {
			rooms = nil
			return
		}
		rooms = append(rooms, *roomStruct)
	}
	return
}

func (r *RoomUsersRepositoryImpl) CreateUserRoomConnection(user model.RoomUsersDto) (*model.RoomUsersDto, error) {
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
