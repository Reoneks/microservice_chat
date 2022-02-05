package user

type userService struct {
	userRep IUserRepository
}

type IUserService interface {
	GetUserByID(id string) (map[string]interface{}, error)
	GetUsers() ([]map[string]interface{}, error)
	CreateUser(user map[string]interface{}) (map[string]interface{}, error)
	UpdateUser(user map[string]interface{}) (map[string]interface{}, error)
	DeleteUser(id string) error
}

func (us *userService) GetUserByID(id string) (map[string]interface{}, error) {
	userDto, err := us.userRep.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return userDto, nil
}

func (us *userService) GetUsers() ([]map[string]interface{}, error) {
	userDto, err := us.userRep.GetUsers()
	if err != nil {
		return nil, err
	}
	return userDto, nil
}

func (us *userService) CreateUser(user map[string]interface{}) (map[string]interface{}, error) {
	userDto, err := us.userRep.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return userDto, nil
}

func (us *userService) UpdateUser(user map[string]interface{}) (map[string]interface{}, error) {
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
