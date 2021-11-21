package auth

import (
	"encoding/json"

	"github.com/Reoneks/microservice_chat/auth/model"
	"github.com/Reoneks/microservice_chat/proto"
)

func ToAuthUser(protoUser *proto.RegistrationRequest) *model.Auth {
	return &model.Auth{
		Email:    protoUser.Email,
		Password: protoUser.Password,
	}
}

func ToUser(protoUser *proto.RegistrationRequest) *proto.UserStruct {
	bytes, _ := json.Marshal(&protoUser)
	return &proto.UserStruct{
		UserInfo: bytes,
	}
}
