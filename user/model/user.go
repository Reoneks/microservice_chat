package model

type User struct {
	ID    string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (User) TableName() string {
	return "users"
}