package database

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type Campaign struct {
	CampaignId   uuid.UUID `gorm:"type:uuid;primaryKey;"`
	CampaignName string
	Description  string
	StartDate    time.Time
	EndDate      time.Time
	*gorm.Model
}
