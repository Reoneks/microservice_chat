package room

import (
	"strings"

	"github.com/Reoneks/microservice_chat/room/model"

	gm "gorm.io/gorm"
)

type roomRepository struct {
	db *gm.DB
}

type IRoomRepository interface {
	GetRoom(id string) (*model.RoomsDto, error)
	CreateRoom(user *model.RoomsDto) (*model.RoomsDto, error)
	UpdateRoom(user *model.RoomsDto) (*model.RoomsDto, error)
	DeleteRoom(id string) error
	GetRooms(filter *RoomsFilter) ([]model.RoomsDto, error)
}

func (r *roomRepository) GetRoom(id string) (*model.RoomsDto, error) {
	room := &model.RoomsDto{}
	if err := r.db.Where("id = ?", id).First(room).Error; err != nil {
		return nil, err
	}
	return room, nil
}

func (r *roomRepository) CreateRoom(room *model.RoomsDto) (*model.RoomsDto, error) {
	if err := r.db.Create(&room).Error; err != nil {
		return nil, err
	}
	return room, nil
}

func (r *roomRepository) UpdateRoom(room *model.RoomsDto) (*model.RoomsDto, error) {
	if err := r.db.Save(&room).Error; err != nil {
		return nil, err
	}
	return room, nil
}

func (r *roomRepository) DeleteRoom(id string) error {
	if err := r.db.Delete(&model.RoomsDto{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *roomRepository) GetRooms(filter *RoomsFilter) (rooms []model.RoomsDto, err error) {
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

func NewRoomRepository(db *gm.DB) IRoomRepository {
	return &roomRepository{
		db,
	}
}
