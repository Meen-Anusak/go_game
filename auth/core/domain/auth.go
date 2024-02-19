package domain

import "github.com/google/uuid"

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	UserID      uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"user_id"`
	UserName    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	FirstName   string    `json:"first_Name"`
	LastName    string    `json:"last_name"`
}
