package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type LeaderBoardEntry struct {
	LeaderBoardEntryID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	PlayerId           uuid.UUID
	Player             Player `gorm:"ForeignKey:PlayerId"`
	CampaignId         uuid.UUID
	Campaign           Campaign `gorm:"ForeignKey:CampaignId"`
	Score              float64
	CreatedAt          time.Time `gorm:"autoTimeCreate;"`
	UpdatedAt          time.Time `gorm:"autoTimeUpdate;"`
	DeletedAt          gorm.DeletedAt
}
