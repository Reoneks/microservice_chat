package repository

import (
	"chatex/messages/messages/messages_interface"
)

func FromMessagesDto(messagesDto messages_interface.MessagesDto) *messages_interface.Message {
	return &messages_interface.Message{
		ID:        messagesDto.ID,
		Text:      messagesDto.Text,
		Status:    messages_interface.NewStatusType(messagesDto.Status),
		RoomID:    messagesDto.RoomID,
		CreatedBy: messagesDto.CreatedBy,
		CreatedAt: messagesDto.CreatedAt,
		UpdatedAt: messagesDto.UpdatedAt,
	}
}

func FromMessagesDtos(messagesDto []messages_interface.MessagesDto) (messages []messages_interface.Message) {
	for _, dto := range messagesDto {
		messages = append(messages, *FromMessagesDto(dto))
	}
	return
}

func ToMessagesDto(messages messages_interface.Message) *messages_interface.MessagesDto {
	return &messages_interface.MessagesDto{
		ID:        messages.ID,
		Text:      messages.Text,
		Status:    messages.Status.ToInt64(),
		RoomID:    messages.RoomID,
		CreatedBy: messages.CreatedBy,
		CreatedAt: messages.CreatedAt,
		UpdatedAt: messages.UpdatedAt,
	}
}

func ToMessagesDtos(messages []messages_interface.Message) (messagesDto []messages_interface.MessagesDto) {
	for _, dto := range messages {
		messagesDto = append(messagesDto, *ToMessagesDto(dto))
	}
	return
}
