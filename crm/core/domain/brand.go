package domain

import (
	"github.com/google/uuid"
	"time"
)

type Brand struct {
	BrandID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"brand_id"`
	BrandName string    `json:"brand_name"`
	CreatedAt time.Time `gorm:"autoTimeCreate;" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoTimeUpdate;" json:"updated_at"`
}
