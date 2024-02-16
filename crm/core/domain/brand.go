package domain

import (
	"github.com/google/uuid"
	"time"
)

type Brand struct {
	BrandID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();"`
	BrandName string
	CreatedAt time.Time `gorm:"autoTimeCreate;"`
	UpdatedAt time.Time `gorm:"autoTimeUpdate;"`
}
