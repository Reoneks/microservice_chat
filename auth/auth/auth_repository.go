package auth

import (
	"github.com/Reoneks/microservice_chat/auth/model"
	"github.com/upper/db/v4"
)

type AuthRepository interface {
	GetUserByEmail(email string) (authUser *model.Auth, err error)
	AddUser(authUser *model.Auth) (*model.Auth, error)
	DeleteUser(id string) error
}

type AuthRepositoryImpl struct {
	session   db.Session
	tableName string
}

func NewAuthRepository(session db.Session) AuthRepository {
	return &AuthRepositoryImpl{
		session:   session,
		tableName: "auth_user",
	}
}

func (auth *AuthRepositoryImpl) GetUserByEmail(email string) (authUser *model.Auth, err error) {
	err = auth.session.SQL().SelectFrom(auth.tableName).Where("email = ?", email).One(&authUser)
	return
}

func (auth *AuthRepositoryImpl) AddUser(authUser *model.Auth) (*model.Auth, error) {
	_, err := auth.session.SQL().InsertInto(auth.tableName).Values(authUser).Exec()
	return authUser, err
}

func (auth *AuthRepositoryImpl) DeleteUser(id string) error {
	_, err := auth.session.SQL().DeleteFrom(auth.tableName).Where("id = ?", id).Exec()
	return err
}
