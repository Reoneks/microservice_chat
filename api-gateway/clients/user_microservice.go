package clients

import (
	"context"
	"net/http"

	"chatex/proto"

	"github.com/labstack/echo/v4"
)

type UserMicroservice struct {
	user proto.UserService
	auth proto.AuthService
}

func (u *UserMicroservice) GetUsers(ctx echo.Context) error {
	rsp, err := u.user.GetUsers(context.Background(), &proto.Empty{})
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !rsp.Status.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Status.Error)
	}
	return ctx.JSON(http.StatusOK, rsp)
}

func (u *UserMicroservice) GetUserByID(ctx echo.Context) error {
	id := ctx.Param("id")

	rsp, err := u.user.GetUserByID(context.Background(), &proto.UserID{UserID: id})
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !rsp.Status.Ok {
		return ctx.String(http.StatusInternalServerError, rsp.Status.Error)
	}
	return ctx.JSON(http.StatusOK, rsp)
}

func (u *UserMicroservice) UpdateUser(ctx echo.Context) error {
	var req proto.UserStruct
	if err := ctx.Bind(&req); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	req.ID = ctx.Get("user_id").(string)

	rsp, err := u.user.UpdateUser(context.Background(), &req)
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

	authRSP, err := u.auth.Delete(context.Background(), &req)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !authRSP.Status.Ok {
		return ctx.JSON(http.StatusInternalServerError, authRSP.Status.Error)
	}

	userRSP, err := u.user.DeleteUser(context.Background(), &req)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if !userRSP.Status.Ok {
		return ctx.JSON(http.StatusInternalServerError, userRSP.Status.Error)
	}

	return ctx.JSON(http.StatusOK, userRSP)
}

func NewUserMicroservice(user proto.UserService, auth proto.AuthService) *UserMicroservice {
	return &UserMicroservice{
		user: user,
		auth: auth,
	}
}
