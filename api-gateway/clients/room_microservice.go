package clients

import (
	"context"
	"net/http"

	"github.com/Reoneks/microservice_chat/proto"

	"github.com/labstack/echo/v4"
)

type RoomMicroservice struct {
	room proto.RoomsService
}

func (u *RoomMicroservice) GetRoom(ctx echo.Context) error {
	rsp, err := u.room.GetRoom(context.Background(), &proto.RoomID{ID: ctx.Param("id")})
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !rsp.Status.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Status.Error)
	}
	return ctx.JSON(http.StatusOK, rsp)
}

func (u *RoomMicroservice) GetRooms(ctx echo.Context) error {
	var filter proto.Filter
	if err := ctx.Bind(&filter); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	rsp, err := u.room.GetRooms(context.Background(), &filter)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !rsp.Status.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Status.Error)
	}
	return ctx.JSON(http.StatusOK, rsp)
}

func (u *RoomMicroservice) CreateRoom(ctx echo.Context) error {
	var room proto.RoomStruct
	if err := ctx.Bind(&room); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	rsp, err := u.room.CreateRoom(context.Background(), &room)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !rsp.Status.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Status.Error)
	}
	return ctx.JSON(http.StatusOK, rsp)
}

func (u *RoomMicroservice) DeleteRoom(ctx echo.Context) error {
	var room proto.DeleteRequest
	if err := ctx.Bind(&room); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	rsp, err := u.room.DeleteRoom(context.Background(), &room)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !rsp.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Error)
	}
	return ctx.JSON(http.StatusOK, rsp)
}

func (u *RoomMicroservice) UpdateRoom(ctx echo.Context) error {
	var room proto.UpdateRequest
	if err := ctx.Bind(&room); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	rsp, err := u.room.UpdateRoom(context.Background(), &room)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !rsp.Status.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Status.Error)
	}
	return ctx.JSON(http.StatusOK, rsp)
}

func (u *RoomMicroservice) AddUsers(ctx echo.Context) error {
	var room proto.AddUsersRequest
	if err := ctx.Bind(&room); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	rsp, err := u.room.AddUsers(context.Background(), &room)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !rsp.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Error)
	}
	return ctx.JSON(http.StatusOK, rsp)
}

func NewRoomMicroservice(room proto.RoomsService) *RoomMicroservice {
	return &RoomMicroservice{
		room: room,
	}
}
