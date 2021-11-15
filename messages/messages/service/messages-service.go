package service

import (
	"chatex/messages/messages/messages_interface"
	"chatex/messages/messages/repository"
)

type messagesService struct {
	messagesService messages_interface.MessagesRepository
}

func (us *messagesService) CreateMessage(message *messages_interface.Message) (*messages_interface.Message, error) {
	userDto, err := us.messagesService.CreateMessage(repository.ToMessagesDto(*message))
	if err != nil {
		return nil, err
	}
	return repository.FromMessagesDto(*userDto), nil
}

func (us *messagesService) UpdateMessage(message *messages_interface.Message) (*messages_interface.Message, error) {
	userDto, err := us.messagesService.UpdateMessage(repository.ToMessagesDto(*message))
	if err != nil {
		return nil, err
	}
	return repository.FromMessagesDto(*userDto), nil
}

func (us *messagesService) DeleteMessage(id string) error {
	return us.messagesService.DeleteMessage(id)
}

func NewMessagesService(messageService messages_interface.MessagesRepository) messages_interface.MessagesService {
	return &messagesService{
		messagesService: messageService,
	}
}
