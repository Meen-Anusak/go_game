package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UserID      uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"user_id"`
	UserName    string    `json:"user_name" validate:"required,min=3"`
	PhoneNumber string    `json:"phone_number" validate:"required,min=10,max=10"`
	Email       string    `json:"email" validate:"required,min=3,email"`
	FirstName   string    `json:"first_name" validate:"required,min=3"`
	LastName    string    `json:"last_name" validate:"required,min=3"`
	BradId      *uuid.UUID
	Brand       *Brand     `gorm:"ForeignKey:BradId" json:"brand"`
	RoleId      *uuid.UUID `json:"role_id"`
	Role        *Role      `gorm:"ForeignKey:RoleId" json:"role"`
	CreatedAt   time.Time  `gorm:"autoTimeCreate;" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoTimeUpdate;" json:"updated_at"`
}
