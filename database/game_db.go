package database

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"gorm.io/datatypes"
)

type Game struct {
	GameId     uuid.UUID `gorm:"type:uuid;primaryKey;"`
	GameName   string
	GameConfig GameConfig
	Campaign   Campaign
	*gorm.Model
}

type GameConfig struct {
	GameConfigId uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Config       datatypes.JSON
}
