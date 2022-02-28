package clients

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Reoneks/microservice_chat/api-gateway/model"
	"github.com/Reoneks/microservice_chat/proto"
	"github.com/asim/go-micro/v3/client"

	"github.com/labstack/echo/v4"
)

type MessagesMicroservice struct {
	messages     proto.MessagesService
	msgAddr      string
	DefaltLimit  int64
	DefaltOffset int64
}

func (u *MessagesMicroservice) GetMessagesByRoom(ctx echo.Context) error {
	limit, err := strconv.ParseInt(ctx.QueryParam("limit"), 10, 64)
	if err != nil {
		limit = u.DefaltLimit
	}

	offset, err := strconv.ParseInt(ctx.QueryParam("offset"), 10, 64)
	if err != nil {
		offset = u.DefaltOffset
	}

	var messageInfo proto.MessageID
	messageInfo.RoomID = ctx.Param("id")
	messageInfo.Limit = limit
	messageInfo.Offset = offset

	rsp, err := u.messages.GetMessagesByRoom(context.Background(), &messageInfo, client.WithAddress(u.msgAddr))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	} else if !rsp.Status.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Status.Error)
	}

	var resp model.PaginationMessagesResponse
	err = json.Unmarshal(rsp.Messages, &resp.Messages)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	resp.Limit = limit
	resp.Offset = offset
	return ctx.JSON(http.StatusOK, resp)
}

func NewMessagesMicroservice(messages proto.MessagesService, DefaltLimit, DefaltOffset int64, msgAddr string) *MessagesMicroservice {
	return &MessagesMicroservice{
		messages:     messages,
		DefaltLimit:  DefaltLimit,
		DefaltOffset: DefaltOffset,
		msgAddr:      msgAddr,
	}
}
