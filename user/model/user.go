package model

type User struct {
	ID    string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name  string
	Email string
}

func (User) TableName() string {
	return "users"
}
