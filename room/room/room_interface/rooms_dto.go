package room

import (
	"time"
)

type RoomsDto struct {
	ID        string    `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	Status    int64     `gorm:"column:status"`
	CreatedBy string    `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (RoomsDto) TableName() string {
	return "rooms"
}
