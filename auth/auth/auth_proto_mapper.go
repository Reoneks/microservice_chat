package auth

import (
	"chatex/proto"
)

func ToAuthUser(protoUser *proto.RegistrationRequest) *AuthUser {
	return &AuthUser{
		Email:    protoUser.Email,
		Password: protoUser.Password,
	}
}

func ToUser(protoUser *proto.RegistrationRequest) *proto.UserStruct {
	return &proto.UserStruct{
		Name:  protoUser.Name,
		Email: protoUser.Email,
	}
}
