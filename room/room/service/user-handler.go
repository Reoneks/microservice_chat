package service

import (
	"chatex/proto"
	"chatex/user/user/user_interface"
	"context"
)

type UserMicro struct {
	UserService user_interface.UserService
}

func (u *UserMicro) GetUserByID(ctx context.Context, req *proto.UserID, rsp *proto.UserStruct) error {
	user, err := u.UserService.GetUserByID(req.UserID)
	if err != nil {
		return err
	}
	ToUserStructRSP(user, rsp)
	return nil
}

func (u *UserMicro) GetUsers(ctx context.Context, req *proto.Empty, rsp *proto.GetUsersResponse) error {
	users, err := u.UserService.GetUsers()
	if err != nil {
		return err
	}
	rsp.Users = ToUserStructs(users)
	return nil
}

func (u *UserMicro) CreateUser(ctx context.Context, req *proto.UserStruct, rsp *proto.UserStruct) error {
	user, err := u.UserService.CreateUser(FromUserStruct(req))
	if err != nil {
		return err
	}
	ToUserStructRSP(user, rsp)
	return nil
}

func (u *UserMicro) UpdateUser(ctx context.Context, req *proto.UserStruct, rsp *proto.UserStruct) error {
	user, err := u.UserService.UpdateUser(FromUserStruct(req))
	if err != nil {
		return err
	}
	ToUserStructRSP(user, rsp)
	return nil
}

func (u *UserMicro) DeleteUser(ctx context.Context, req *proto.UserID, rsp *proto.DeleteUserResponse) error {
	err := u.UserService.DeleteUser(req.UserID)
	if err != nil {
		rsp.Result = "Error"
		rsp.Error = err.Error()
		return err
	}
	rsp.Result = "OK"
	return nil
}
