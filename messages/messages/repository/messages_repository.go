package repository

import (
	"chatex/messages/messages/messages_interface"
	"time"

	gm "gorm.io/gorm"
)

type MessagesRepositoryImpl struct {
	db *gm.DB
}

func (r *MessagesRepositoryImpl) CreateMessage(message *messages_interface.MessagesDto) (*messages_interface.MessagesDto, error) {
	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()
	if err := r.db.Create(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func (r *MessagesRepositoryImpl) UpdateMessage(message *messages_interface.MessagesDto) (*messages_interface.MessagesDto, error) {
	var messages *messages_interface.MessagesDto
	if err := r.db.Where("id = ?", message.ID).First(&messages).Error; err != nil {
		return nil, err
	}
	message.CreatedAt = messages.CreatedAt
	message.UpdatedAt = time.Now()
	if err := r.db.Save(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func (r *MessagesRepositoryImpl) DeleteMessage(id string) error {
	if err := r.db.Delete(&messages_interface.MessagesDto{}, id).Error; err != nil {
		return err
	}
	return nil
}

func NewMessagesRepository(db *gm.DB) messages_interface.MessagesRepository {
	return &MessagesRepositoryImpl{
		db,
	}
}
