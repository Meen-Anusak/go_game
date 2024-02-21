package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UserID      uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"user_id"`
	UserName    string    `json:"user_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	BradId      *uuid.UUID
	Brand       *Brand     `gorm:"ForeignKey:BradId" json:"brand"`
	RoleId      *uuid.UUID `json:"role_id"`
	Role        *Role      `gorm:"ForeignKey:RoleId" json:"role"`
	CreatedAt   time.Time  `gorm:"autoTimeCreate;" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoTimeUpdate;" json:"updated_at"`
}
