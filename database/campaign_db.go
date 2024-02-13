package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Campaign struct {
	CampaignID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();"`
	CampaignName string
	Description  string
	StartDate    time.Time
	EndDate      time.Time
	CreatedAt    time.Time `gorm:"autoTimeCreate;"`
	UpdatedAt    time.Time `gorm:"autoTimeUpdate;"`
	DeletedAt    gorm.DeletedAt
}
