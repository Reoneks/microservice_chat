package auth

import (
	"github.com/Reoneks/microservice_chat/proto"
)

func ToUser(protoUser *proto.RegistrationRequest) *proto.UserStruct {
	return &proto.UserStruct{
		UserInfo: protoUser.UserInfo,
	}
}
