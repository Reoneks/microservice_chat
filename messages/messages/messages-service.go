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
	return us.messagesService.GetMessagesByRoom(roomID, limit, offset)
}

func (us *messagesService) CreateMessage(message map[string]interface{}) (map[string]interface{}, error) {
	return us.messagesService.CreateMessage(message)
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
