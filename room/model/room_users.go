package model

type RoomUsersDto struct {
	UserID string `gorm:"column:user_id"`
	RoomID string `gorm:"column:room_id"`
}

func (RoomUsersDto) TableName() string {
	return "room_users"
}
