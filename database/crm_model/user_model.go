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
	BradId      uuid.UUID
	Brand       Brand `gorm:"ForeignKey:BradId"`
	RoleId      uuid.UUID
	Role        Role      `gorm:"ForeignKey:RoleId"`
	CreatedAt   time.Time `gorm:"autoTimeCreate;"`
	UpdatedAt   time.Time `gorm:"autoTimeUpdate;"`

	DeletedAt gorm.DeletedAt
}
