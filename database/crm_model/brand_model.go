package crm_model

import (
	"time"

	"github.com/google/uuid"
)

type Brand struct {
	BrandID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();"`
	BrandName string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoTimeCreate;"`
	UpdatedAt time.Time `gorm:"autoTimeUpdate;"`
}
