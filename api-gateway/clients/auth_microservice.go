package clients

import (
	"context"
	"net/http"

	"github.com/Reoneks/microservice_chat/proto"
	"github.com/asim/go-micro/v3/client"

	"github.com/labstack/echo/v4"
)

type AuthMicroservice struct {
	auth     proto.AuthService
	authAddr string
}

func NewAuthMicroservice(auth proto.AuthService, authAddr string) *AuthMicroservice {
	return &AuthMicroservice{
		auth:     auth,
		authAddr: authAddr,
	}
}

func (a *AuthMicroservice) Register(ctx echo.Context) error {
	var req proto.RegistrationRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	rsp, err := a.auth.Registration(context.Background(), &req, client.WithAddress(a.authAddr))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	} else if rsp.Token == "" {
		return ctx.String(http.StatusUnauthorized, rsp.Status.Error)
	}

	return ctx.JSON(http.StatusOK, rsp.Token)
}

func (a *AuthMicroservice) Login(ctx echo.Context) error {
	var req proto.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	rsp, err := a.auth.LoginHandler(context.Background(), &req, client.WithAddress(a.authAddr))
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	} else if rsp.Token == "" {
		return ctx.NoContent(http.StatusUnauthorized)
	}

	return ctx.JSON(http.StatusOK, rsp.Token)
}
