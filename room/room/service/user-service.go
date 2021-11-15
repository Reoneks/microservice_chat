package service

import (
	"chatex/user/user/repository"
	"chatex/user/user/user_interface"
)

type userService struct {
	userRep user_interface.UserRepository
}

func (us *userService) GetUserByID(id string) (*user_interface.User, error) {
	userDto, err := us.userRep.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return repository.FromUserDTO(userDto), nil
}

func (us *userService) GetUsers() ([]user_interface.User, error) {
	userDto, err := us.userRep.GetUsers()
	if err != nil {
		return nil, err
	}
	return repository.FromUserDTOs(userDto), nil
}

func (us *userService) CreateUser(user *user_interface.User) (*user_interface.User, error) {
	userDto, err := us.userRep.CreateUser(repository.ToUserDTO(user))
	if err != nil {
		return nil, err
	}
	return repository.FromUserDTO(userDto), nil
}

func (us *userService) UpdateUser(user *user_interface.User) (*user_interface.User, error) {
	userDto, err := us.userRep.UpdateUser(repository.ToUserDTO(user))
	if err != nil {
		return nil, err
	}
	return repository.FromUserDTO(userDto), nil
}

func (us *userService) DeleteUser(id string) error {
	return us.userRep.DeleteUser(id)
}

func NewUserService(userRep user_interface.UserRepository) user_interface.UserService {
	return &userService{
		userRep: userRep,
	}
}
