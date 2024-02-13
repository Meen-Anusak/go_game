package crm_model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	ActivityAT   time.Time
	AuthToken    string
	PlatFrom     string
	CreatedAt    time.Time `gorm:"autoTimeCreate;"`
	UpdatedAt    time.Time `gorm:"autoTimeUpdate;"`
	DeletedAt    gorm.DeletedAt
}
