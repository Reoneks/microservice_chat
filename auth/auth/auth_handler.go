package auth

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/Reoneks/microservice_chat/auth/utils"
	"github.com/Reoneks/microservice_chat/proto"

	"github.com/go-chi/jwtauth"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=4,max=32,alphanum"`
}

type Auth struct {
	authService AuthService
	userService proto.UserService
	jwt         *jwtauth.JWTAuth
}

func NewAuth(authService AuthService, userService proto.UserService, jwt *jwtauth.JWTAuth) *Auth {
	return &Auth{
		authService: authService,
		userService: userService,
		jwt:         jwt,
	}
}

func (a *Auth) AuthHandler(ctx context.Context, req *proto.Token, resp *proto.UserID) error {
	if req.Token == "" {
		resp.Status.Ok = false
		resp.Status.Error = "token is required"
		return errors.New("token is required")
	}

	payload, err := jwtauth.VerifyToken(a.jwt, req.Token)
	if err != nil {
		resp.Status.Ok = false
		resp.Status.Error = err.Error()
		return err
	}

	userID, ok := payload.Get("user_id")
	if !ok {
		resp.Status.Ok = false
		resp.Status.Error = "there is no user id"
		return errors.New("there is no user id")
	}

	resp.UserID = userID.(string)
	resp.Status.Ok = true
	return nil
}

func (a *Auth) Delete(ctx context.Context, req *proto.UserID, resp *proto.DeleteUserResponse) (err error) {
	err = a.authService.Delete(req.UserID)
	if err != nil {
		resp.Status.Ok = false
		resp.Status.Error = err.Error()
		return err
	}

	resp.Status.Ok = true
	return nil
}

func (a *Auth) LoginHandler(ctx context.Context, req *proto.LoginRequest, resp *proto.Token) (err error) {
	resp.Token, err = a.authService.Login(req.Email, req.Password)
	if err != nil {
		resp.Status.Ok = false
		resp.Status.Error = err.Error()
		return err
	} else if resp.Token == "" {
		resp.Status.Ok = false
		resp.Status.Error = "token is empty"
		return errors.New("token is empty")
	}

	resp.Status.Ok = true
	return nil
}

func (a *Auth) Registration(ctx context.Context, req *proto.RegistrationRequest, resp *proto.Token) (err error) {
	userResp, err := a.userService.CreateUser(context.Background(), ToUser(req))
	if err != nil {
		resp.Status.Ok = false
		resp.Status.Error = err.Error()
		return err
	}

	var userRespMap map[string]interface{}
	err = json.Unmarshal(userResp.UserInfo, &userRespMap)
	if err != nil {
		resp.Status.Ok = false
		resp.Status.Error = err.Error()
		return err
	}

	user := make(map[string]interface{})
	user["_id"] = userRespMap["_id"].(string)
	user["email"] = req.Email
	user["password"], err = utils.Encrypt(req.Password)
	if err != nil {
		resp.Status.Ok = false
		resp.Status.Error = err.Error()
		return err
	}

	resp.Token, err = a.authService.Register(user)
	if err != nil {
		resp.Status.Ok = false
		resp.Status.Error = err.Error()
		return err
	} else if resp.Token == "" {
		resp.Status.Ok = false
		resp.Status.Error = "token is empty"
		return errors.New("token is empty")
	}

	resp.Status.Ok = true
	return nil
}
