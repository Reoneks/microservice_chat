package messages_interface

type MessagesRepository interface {
	CreateMessage(message *MessagesDto) (*MessagesDto, error)
	UpdateMessage(message *MessagesDto) (*MessagesDto, error)
	DeleteMessage(id string) error
}

type MessagesService interface {
	CreateMessage(message *Message) (*Message, error)
	UpdateMessage(message *Message) (*Message, error)
	DeleteMessage(id string) error
}
