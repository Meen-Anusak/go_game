package database

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Player struct {
	PlayerId   uuid.UUID `gorm:"type:uuid;primaryKey;"`
	PlayerName string    `gorm:"size:15"`
	*gorm.Model
}

type PlayerRewardLog struct {
	PlayerRewardLogId uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Player            Player
	Reward            Reward
	Campaign          Campaign
	*gorm.Model
}

type PlayerActivityGameLog struct {
	PlayerActivityGameLogId uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Player                  Player
	Game                    Game
	Campaign                Campaign
	Point                   float64
	Stamina                 int64
	*gorm.Model
}
