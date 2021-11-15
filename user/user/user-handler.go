package user

import (
	"context"
	"encoding/json"

	"github.com/Reoneks/microservice_chat/proto"
)

type UserMicro struct {
	UserService IUserService
}

func (u *UserMicro) GetUserByID(ctx context.Context, req *proto.UserID, rsp *proto.UserStruct) error {
	user, err := u.UserService.GetUserByID(req.UserID)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	rsp.UserInfo = bytes
	rsp.Status.Ok = true
	return nil
}

func (u *UserMicro) GetUsers(ctx context.Context, req *proto.Empty, rsp *proto.GetUsersResponse) error {
	users, err := u.UserService.GetUsers()
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	var resp []*proto.UserStruct

	for _, user := range users {
		bytes, err := json.Marshal(user)
		if err != nil {
			rsp.Status.Ok = false
			rsp.Status.Error = err.Error()
			return err
		}

		resp = append(resp, &proto.UserStruct{
			UserInfo: bytes,
		})
	}

	rsp.Users = resp
	rsp.Status.Ok = true
	return nil
}

func (u *UserMicro) CreateUser(ctx context.Context, req *proto.UserStruct, rsp *proto.UserStruct) error {
	var data map[string]interface{}
	err := json.Unmarshal(req.UserInfo, &data)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	user, err := u.UserService.CreateUser(data)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	rsp.UserInfo = bytes
	rsp.Status.Ok = true
	return nil
}

func (u *UserMicro) UpdateUser(ctx context.Context, req *proto.UserStruct, rsp *proto.UserStruct) error {
	var data map[string]interface{}
	err := json.Unmarshal(req.UserInfo, &data)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	user, err := u.UserService.UpdateUser(data)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	rsp.UserInfo = bytes
	rsp.Status.Ok = true
	return nil
}

func (u *UserMicro) DeleteUser(ctx context.Context, req *proto.UserID, rsp *proto.DeleteUserResponse) error {
	err := u.UserService.DeleteUser(req.UserID)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}
	rsp.Status.Ok = true
	return nil
}
