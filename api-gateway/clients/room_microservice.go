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

type RoomMicroservice struct {
	room         proto.RoomsService
	roomAddr     string
	DefaltLimit  int64
	DefaltOffset int64
}

func (u *RoomMicroservice) GetRooms(ctx echo.Context) error {
	limit, err := strconv.ParseInt(ctx.QueryParam("limit"), 10, 64)
	if err != nil {
		limit = u.DefaltLimit
	}

	offset, err := strconv.ParseInt(ctx.QueryParam("offset"), 10, 64)
	if err != nil {
		offset = u.DefaltOffset
	}

	rsp, err := u.room.GetAllRooms(context.Background(), &proto.GetAllRoomsRequest{
		Limit:  limit,
		Offset: offset,
		UserID: ctx.Get("user_id").(string),
	}, client.WithAddress(u.roomAddr))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	} else if !rsp.Status.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Status.Error)
	}

	var resp model.PaginationRoomsResponse
	err = json.Unmarshal(rsp.Room, &resp.Rooms)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	resp.Limit = limit
	resp.Offset = offset

	return ctx.JSON(http.StatusOK, resp)
}

func (u *RoomMicroservice) CreateRoom(ctx echo.Context) error {
	var room map[string]interface{}
	if err := ctx.Bind(&room); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	room["created_by"] = ctx.Get("user_id").(string)
	var (
		req proto.RoomStruct
		err error
	)

	req.RoomInfo, err = json.Marshal(&room)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	rsp, err := u.room.CreateRoom(context.Background(), &req, client.WithAddress(u.roomAddr))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	} else if !rsp.Status.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Status.Error)
	}

	var resp map[string]interface{}
	err = json.Unmarshal(rsp.Room, &resp)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (u *RoomMicroservice) UpdateRoom(ctx echo.Context) error {
	var room map[string]interface{}
	if err := ctx.Bind(&room); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	roomInfo, err := json.Marshal(room)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	var data proto.UpdateRequest
	data.Room = roomInfo
	data.RoomID = ctx.Param("id")

	rsp, err := u.room.UpdateRoom(context.Background(), &data, client.WithAddress(u.roomAddr))
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !rsp.Status.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Status.Error)
	}

	var resp map[string]interface{}
	err = json.Unmarshal(rsp.Room, &resp)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (u *RoomMicroservice) DeleteRoom(ctx echo.Context) error {
	var room proto.DeleteRequest
	room.RoomID = ctx.Param("id")

	rsp, err := u.room.DeleteRoom(context.Background(), &room, client.WithAddress(u.roomAddr))
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !rsp.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Error)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (u *RoomMicroservice) AddUsers(ctx echo.Context) error {
	var userIDs model.UserIDsRequest
	if err := ctx.Bind(&userIDs); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	var req proto.AddUsersRequest
	req.UserIDs = userIDs.UserIDs
	req.RoomID = ctx.Param("id")

	rsp, err := u.room.AddUsers(context.Background(), &req, client.WithAddress(u.roomAddr))
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !rsp.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Error)
	}

	return ctx.NoContent(http.StatusOK)
}

func (u *RoomMicroservice) DeleteUsers(ctx echo.Context) error {
	var userIDs model.UserIDsRequest
	if err := ctx.Bind(&userIDs); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	var req proto.AddUsersRequest
	req.UserIDs = userIDs.UserIDs
	req.RoomID = ctx.Param("id")

	rsp, err := u.room.DeleteUsers(context.Background(), &req, client.WithAddress(u.roomAddr))
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !rsp.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Error)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func NewRoomMicroservice(room proto.RoomsService, DefaltLimit, DefaltOffset int64, roomAddr string) *RoomMicroservice {
	return &RoomMicroservice{
		room:         room,
		DefaltLimit:  DefaltLimit,
		DefaltOffset: DefaltOffset,
		roomAddr:     roomAddr,
	}
}
