package messages

import "github.com/Reoneks/microservice_chat/messages/model"

type messagesService struct {
	messagesService MessagesRepository
}

type MessagesService interface {
	CreateMessage(message *model.Message) (*model.Message, error)
	UpdateMessage(message *model.Message) (*model.Message, error)
	DeleteMessage(id string) error
}

func (us *messagesService) CreateMessage(message *model.Message) (*model.Message, error) {
	userDto, err := us.messagesService.CreateMessage(message)
	if err != nil {
		return nil, err
	}
	return userDto, nil
}

func (us *messagesService) UpdateMessage(message *model.Message) (*model.Message, error) {
	upMessage, err := us.messagesService.UpdateMessage(message)
	if err != nil {
		return nil, err
	}
	return upMessage, nil
}

func (us *messagesService) DeleteMessage(id string) error {
	return us.messagesService.DeleteMessage(id)
}

func NewMessagesService(messageService MessagesRepository) MessagesService {
	return &messagesService{
		messagesService: messageService,
	}
}
