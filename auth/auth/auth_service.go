package auth

import (
	"github.com/Reoneks/microservice_chat/auth/utils"

	"github.com/go-chi/jwtauth"
)

type AuthService interface {
	Login(email, password string) (string, error)
	Register(user map[string]interface{}) (string, error)
	Delete(id string) error
}

type AuthServiceImpl struct {
	jwt            *jwtauth.JWTAuth
	authRepository AuthRepository
}

func (s *AuthServiceImpl) Login(email, password string) (string, error) {
	receivedUser, err := s.authRepository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	err = utils.Compare(receivedUser["password"].(string), password)
	if err != nil {
		return "", err
	}
	_, tokenString, err := s.jwt.Encode(map[string]interface{}{"user_id": receivedUser["_id"]})
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *AuthServiceImpl) Register(user map[string]interface{}) (string, error) {
	receivedUser, err := s.authRepository.AddUser(user)
	if err != nil {
		return "", err
	}

	_, tokenString, err := s.jwt.Encode(map[string]interface{}{"user_id": receivedUser["_id"]})
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthServiceImpl) Delete(id string) error {
	return s.authRepository.DeleteUser(id)
}

func NewAuthService(authRepository AuthRepository, jwt *jwtauth.JWTAuth) AuthService {
	return &AuthServiceImpl{
		jwt,
		authRepository,
	}
}
