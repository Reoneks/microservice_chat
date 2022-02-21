package messages

type messagesService struct {
	messagesService MessagesRepository
}

type MessagesService interface {
	GetMessagesByRoom(roomID string, limit, offset int64) ([]map[string]interface{}, error)
	CreateMessage(message map[string]interface{}) (map[string]interface{}, error)
	UpdateMessage(message map[string]interface{}) (map[string]interface{}, error)
	DeleteMessage(id string) error
}

func (us *messagesService) GetMessagesByRoom(roomID string, limit, offset int64) ([]map[string]interface{}, error) {
	messages, err := us.messagesService.GetMessagesByRoom(roomID, limit, offset)
	if err != nil {
		return nil, err
	}

	for i := range messages {
		messages[i]["id"] = messages[i]["_id"]
		delete(messages[i], "_id")
	}

	return messages, nil
}

func (us *messagesService) CreateMessage(message map[string]interface{}) (map[string]interface{}, error) {
	result, err := us.messagesService.CreateMessage(message)
	if err != nil {
		return nil, err
	}

	result["id"] = result["_id"]
	delete(result, "_id")
	return result, nil
}

func (us *messagesService) UpdateMessage(message map[string]interface{}) (map[string]interface{}, error) {
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
