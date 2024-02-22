package domain

import (
	"github.com/google/uuid"
	"time"
)

type Customer struct {
	CustomerID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();"`
	CustomerName string
	FistName     string
	LastName     string
	Email        string
	PhoneNumber  string
	BrandId      uuid.UUID
	Brand        Brand `gorm:"ForeignKey:BrandId"`
	ActivityAt   time.Time
	AuthToken    string
	PlatFrom     string
	CreatedAt    time.Time `gorm:"autoTimeCreate;"`
	UpdatedAt    time.Time `gorm:"autoTimeUpdate;"`
}
type CustomerRequestSchema struct {
	CustomerName string `json:"customer_name" validate:"required,min=3"`
	FistName     string `json:"fist_name" validate:"required,min=3"`
	LastName     string `json:"last_name" validate:"required,min=3"`
	Email        string `json:"email" validate:"required,min=3,email"`
	PhoneNumber  string `json:"phone_number" validate:"required,min=10,max=10"`
}

type CustomerResponseSchema struct {
	CustomerName string    `json:"customer_name" validate:"required,min=3"`
	FistName     string    `json:"fist_name" validate:"required,min=3"`
	LastName     string    `json:"last_name" validate:"required,min=3"`
	Email        string    `json:"email" validate:"required,min=3,email"`
	PhoneNumber  string    `json:"phone_number" validate:"required,min=10,max=10"`
	CreatedAt    time.Time `gorm:"autoTimeCreate;"`
	UpdatedAt    time.Time `gorm:"autoTimeUpdate;"`
}
