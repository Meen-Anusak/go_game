package game_model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Game struct {
	GameID       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	GameName     string
	GameConfigId uuid.UUID
	GameConfig   GameConfig `gorm:"ForeignKey:GameConfigId"`
	CreatedAt    time.Time  `gorm:"autoTimeCreate;"`
	UpdatedAt    time.Time  `gorm:"autoTimeUpdate;"`
	DeletedAt    gorm.DeletedAt
	// Campaign   Campaign
}

type GameConfig struct {
	GameConfigID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	// GameID       string
	// Game         Game `gorm:"ForeignKey:GameConfigID"`
	Config    datatypes.JSON
	CreatedAt time.Time `gorm:"autoTimeCreate;"`
	UpdatedAt time.Time `gorm:"autoTimeUpdate;"`
	DeletedAt gorm.DeletedAt
}
