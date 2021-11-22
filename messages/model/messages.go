package model

import "time"

type StatusType int64

//^ Status types
const (
	Unread StatusType = 1
	Read   StatusType = 2
)

type Message struct {
	ID        string     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Text      string     `json:"text"`
	Status    StatusType `json:"status"`
	RoomID    string     `json:"room_id"`
	CreatedBy string     `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	MessageType string `gorm:"-" json:"message_type"`
}