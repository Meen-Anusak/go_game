package database

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type LeaderBoardEntry struct {
	LeaderBoardEntryID uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Player             Player
	Campaign           Campaign
	Score              float64
	*gorm.Model
}
