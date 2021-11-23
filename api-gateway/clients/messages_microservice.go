package clients

import (
	"context"
	"net/http"

	"github.com/Reoneks/microservice_chat/proto"

	"github.com/labstack/echo/v4"
)

type MessagesMicroservice struct {
	messages proto.MessagesService
}

func (u *MessagesMicroservice) GetMessagesByRoom(ctx echo.Context) error {
	var messageInfo proto.MessageID
	if err := ctx.Bind(&messageInfo); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	rsp, err := u.messages.GetMessagesByRoom(context.Background(), &messageInfo)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !rsp.Status.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Status.Error)
	}
	return ctx.JSON(http.StatusOK, rsp)
}

func NewMessagesMicroservice(messages proto.MessagesService) *MessagesMicroservice {
	return &MessagesMicroservice{
		messages: messages,
	}
}
