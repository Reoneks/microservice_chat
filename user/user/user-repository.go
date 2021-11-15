package user

import (
	"user_service/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	GetUserByID(id string) (map[string]interface{}, error)
	GetUsers() ([]map[string]interface{}, error)
	CreateUser(user map[string]interface{}) (map[string]interface{}, error)
	UpdateUser(user map[string]interface{}) (map[string]interface{}, error)
	DeleteUser(id string) error
}

func (ur *userRepository) GetUserByID(id string) (result map[string]interface{}, err error) {
	err = ur.db.First(&result, "id = ?", id).Error
	return
}

func (ur *userRepository) GetUsers() (result []map[string]interface{}, err error) {
	err = ur.db.Find(&result).Error
	return
}

func (ur *userRepository) CreateUser(user map[string]interface{}) (map[string]interface{}, error) {
	err := ur.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) UpdateUser(user map[string]interface{}) (map[string]interface{}, error) {
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
