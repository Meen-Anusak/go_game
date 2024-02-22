package domain

import (
	"github.com/google/uuid"
	"time"
)

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
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

type Auth struct {
	UserName    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Token       string    `json:"token"`
	Expired     time.Time `json:"expired"`
}
