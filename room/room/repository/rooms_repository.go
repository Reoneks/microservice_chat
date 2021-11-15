package repository

import (
	room "chatex/room/room/room_interface"
	"strings"
	"time"

	gm "gorm.io/gorm"
)

type RoomRepositoryImpl struct {
	db *gm.DB
}

func (r *RoomRepositoryImpl) GetRoom(id string) (*room.RoomsDto, error) {
	room := &room.RoomsDto{}
	if err := r.db.Where("id = ?", id).First(room).Error; err != nil {
		return nil, err
	}
	return room, nil
}

func (r *RoomRepositoryImpl) CreateRoom(room room.RoomsDto) (*room.RoomsDto, error) {
	room.CreatedAt = time.Now()
	if err := r.db.Create(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepositoryImpl) UpdateRoom(room room.RoomsDto) (*room.RoomsDto, error) {
	result, err := r.GetRoom(room.ID)
	if err != nil {
		return nil, err
	}
	room.CreatedAt = result.CreatedAt
	if err := r.db.Save(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepositoryImpl) DeleteRoom(id string) error {
	if err := r.db.Delete(&room.RoomsDto{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *RoomRepositoryImpl) GetRooms(filter *room.RoomsFilter) (rooms []room.RoomsDto, err error) {
	var findResult *gm.DB = r.db
	var search []string
	if filter != nil {
		if len(filter.IDs) > 0 {
			search = append(search, "id IN ("+strings.Join(filter.IDs, ",")+")")
		}
		if len(search) > 0 {
			findResult = findResult.Where(strings.Join(search, " AND "))
		}
	}
	if err := findResult.Find(&rooms).Error; err != nil {
		return nil, err
	}
	return
}

func NewRoomRepository(db *gm.DB) room.RoomRepository {
	return &RoomRepositoryImpl{
		db,
	}
}
