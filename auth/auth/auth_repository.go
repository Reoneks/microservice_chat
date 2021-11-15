package auth

import "github.com/upper/db/v4"

type AuthRepository interface {
	GetUserByEmail(email string) (authUser *AuthUserDTO, err error)
	AddUser(authUser *AuthUserDTO) (*AuthUserDTO, error)
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

func (auth *AuthRepositoryImpl) GetUserByEmail(email string) (authUser *AuthUserDTO, err error) {
	err = auth.session.SQL().SelectFrom(auth.tableName).Where("email = ?", email).One(&authUser)
	return
}

func (auth *AuthRepositoryImpl) AddUser(authUser *AuthUserDTO) (*AuthUserDTO, error) {
	_, err := auth.session.SQL().InsertInto(auth.tableName).Values(authUser).Exec()
	return authUser, err
}

func (auth *AuthRepositoryImpl) DeleteUser(id string) error {
	_, err := auth.session.SQL().DeleteFrom(auth.tableName).Where("id = ?", id).Exec()
	return err
}
