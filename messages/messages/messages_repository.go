package messages

import (
	"time"

	"github.com/Reoneks/microservice_chat/messages/model"
	gm "gorm.io/gorm"
)

type MessagesRepository interface {
	CreateMessage(message *model.Message) (*model.Message, error)
	UpdateMessage(message *model.Message) (*model.Message, error)
	DeleteMessage(id string) error
}

type MessagesRepositoryImpl struct {
	db *gm.DB
}

func (r *MessagesRepositoryImpl) CreateMessage(message *model.Message) (*model.Message, error) {
	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()
	if err := r.db.Create(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func (r *MessagesRepositoryImpl) UpdateMessage(message *model.Message) (*model.Message, error) {
	var messages *model.Message
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
	if err := r.db.Delete(&model.Message{}, id).Error; err != nil {
		return err
	}
	return nil
}

func NewMessagesRepository(db *gm.DB) MessagesRepository {
	return &MessagesRepositoryImpl{
		db,
	}
}
