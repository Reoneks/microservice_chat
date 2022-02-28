package clients

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Reoneks/microservice_chat/proto"
	"github.com/asim/go-micro/v3/client"

	"github.com/labstack/echo/v4"
)

type UserMicroservice struct {
	user proto.UserService
	auth proto.AuthService

	userAddr string
	authAddr string
}

func (u *UserMicroservice) GetUsers(ctx echo.Context) error {
	rsp, err := u.user.GetUsers(context.Background(), &proto.Empty{}, client.WithAddress(u.userAddr))
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !rsp.Status.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Status.Error)
	}
	return ctx.JSON(http.StatusOK, rsp)
}

func (u *UserMicroservice) GetUserByID(ctx echo.Context) error {
	id := ctx.Param("id")

	rsp, err := u.user.GetUserByID(context.Background(), &proto.UserID{UserID: id}, client.WithAddress(u.userAddr))
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !rsp.Status.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Status.Error)
	}
	return ctx.JSON(http.StatusOK, rsp)
}

func (u *UserMicroservice) UpdateUser(ctx echo.Context) error {
	var req proto.UserStruct
	var user map[string]interface{}
	var err error

	if err = ctx.Bind(&user); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	user["_id"] = ctx.Get("user_id").(string)
	req.UserInfo, err = json.Marshal(user)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	rsp, err := u.user.UpdateUser(context.Background(), &req, client.WithAddress(u.userAddr))
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !rsp.Status.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Status.Error)
	}
	return ctx.JSON(http.StatusOK, rsp)
}

func (u *UserMicroservice) DeleteUser(ctx echo.Context) error {

	ID := ctx.Get("user_id").(string)
	req := proto.UserID{UserID: ID}

	authRSP, err := u.auth.Delete(context.Background(), &req, client.WithAddress(u.authAddr))
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !authRSP.Status.Ok {
		return ctx.JSON(http.StatusInternalServerError, authRSP.Status.Error)
	}

	userRSP, err := u.user.DeleteUser(context.Background(), &req, client.WithAddress(u.userAddr))
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !userRSP.Status.Ok {
		return ctx.JSON(http.StatusInternalServerError, userRSP.Status.Error)
	}

	return ctx.JSON(http.StatusOK, userRSP)
}

func NewUserMicroservice(user proto.UserService, auth proto.AuthService, userAddr, authAddr string) *UserMicroservice {
	return &UserMicroservice{
		user:     user,
		auth:     auth,
		userAddr: userAddr,
		authAddr: authAddr,
	}
}
