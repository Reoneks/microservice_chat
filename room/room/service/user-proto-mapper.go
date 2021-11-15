package service

import (
	"chatex/proto"
	"chatex/user/user/user_interface"
)

func FromUserStruct(userStruct *proto.UserStruct) *user_interface.User {
	return &user_interface.User{
		ID:    userStruct.ID,
		Name:  userStruct.Name,
		Email: userStruct.Email,
	}
}

func FromUserStructs(userStructs []*proto.UserStruct) (users []user_interface.User) {
	for _, userStruct := range userStructs {
		users = append(users, *FromUserStruct(userStruct))
	}
	return
}

func ToUserStruct(user *user_interface.User) *proto.UserStruct {
	return &proto.UserStruct{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToUserStructRSP(user *user_interface.User, rsp *proto.UserStruct) {
	rsp.ID = user.ID
	rsp.Name = user.Name
	rsp.Email = user.Email
}

func ToUserStructs(users []user_interface.User) (UserStructs []*proto.UserStruct) {
	for _, dto := range users {
		UserStructs = append(UserStructs, ToUserStruct(&dto))
	}
	return
}
