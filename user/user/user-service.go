package user

import "github.com/Reoneks/microservice_chat/user/model"

type userService struct {
	userRep IUserRepository
}

type IUserService interface {
	GetUserByID(id string) (*model.User, error)
	GetUsers() ([]model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(id string) error
}

func (us *userService) GetUserByID(id string) (*model.User, error) {
	userDto, err := us.userRep.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return userDto, nil
}

func (us *userService) GetUsers() ([]model.User, error) {
	userDto, err := us.userRep.GetUsers()
	if err != nil {
		return nil, err
	}
	return userDto, nil
}

func (us *userService) CreateUser(user *model.User) (*model.User, error) {
	userDto, err := us.userRep.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return userDto, nil
}

func (us *userService) UpdateUser(user *model.User) (*model.User, error) {
	userDto, err := us.userRep.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return userDto, nil
}

func (us *userService) DeleteUser(id string) error {
	return us.userRep.DeleteUser(id)
}

func NewUserService(userRep IUserRepository) IUserService {
	return &userService{
		userRep: userRep,
	}
}
