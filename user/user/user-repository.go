package user

import (
	"github.com/Reoneks/microservice_chat/user/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	GetUserByID(id string) (*model.User, error)
	GetUsers() ([]model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(id string) error
}

func (ur *userRepository) GetUserByID(id string) (result *model.User, err error) {
	err = ur.db.First(&result, "id = ?", id).Error
	return
}

func (ur *userRepository) GetUsers() (result []model.User, err error) {
	err = ur.db.Find(&result).Error
	return
}

func (ur *userRepository) CreateUser(user *model.User) (*model.User, error) {
	err := ur.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) UpdateUser(user *model.User) (*model.User, error) {
	err := ur.db.Model(&model.User{}).Updates(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) DeleteUser(id string) error {
	return ur.db.Delete(&model.User{}, id).Error
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}
