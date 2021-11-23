package messages

import "github.com/Reoneks/microservice_chat/messages/model"

type messagesService struct {
	messagesService MessagesRepository
}

type MessagesService interface {
	GetMessagesByRoom(roomID string, limit, offset int) ([]model.Message, error)
	CreateMessage(message *model.Message) (*model.Message, error)
	UpdateMessage(message *model.Message) (*model.Message, error)
	DeleteMessage(id string) error
}

func (us *messagesService) GetMessagesByRoom(roomID string, limit, offset int) ([]model.Message, error) {
	return us.messagesService.GetMessagesByRoom(roomID, limit, offset)
}

func (us *messagesService) CreateMessage(message *model.Message) (*model.Message, error) {
	return us.messagesService.CreateMessage(message)
}

func (us *messagesService) UpdateMessage(message *model.Message) (*model.Message, error) {
	return us.messagesService.UpdateMessage(message)
}

func (us *messagesService) DeleteMessage(id string) error {
	return us.messagesService.DeleteMessage(id)
}

func NewMessagesService(messageService MessagesRepository) MessagesService {
	return &messagesService{
		messagesService: messageService,
	}
}
