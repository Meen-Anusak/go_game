package crm_model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserID      uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();"`
	UserName    string
	Password    string
	PhoneNumber string
	Email       string
	FirstName   string
	LastName    string
	BradId      *uuid.UUID
	Brand       *Brand `gorm:"ForeignKey:BradId"`
	RoleId      *uuid.UUID
	Role        *Role     `gorm:"ForeignKey:RoleId"`
	CreatedAt   time.Time `gorm:"autoTimeCreate;"`
	UpdatedAt   time.Time `gorm:"autoTimeUpdate;"`
	DeletedAt   gorm.DeletedAt
}

type UserResponse struct {
	UserID      uuid.UUID `json:"user_id"`
	UserName    string    `json:"user_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Role        *Role     `json:"role"`
}

func FilterUserRecord(user *User) UserResponse {
	return UserResponse{
		UserID:      user.UserID,
		UserName:    user.UserName,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Role:        user.Role,
	}
}
