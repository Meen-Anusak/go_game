package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Player struct {
	PlayerID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	PlayerName string    `gorm:"size:15"`
	CreatedAt  time.Time `gorm:"autoTimeCreate;"`
	UpdatedAt  time.Time `gorm:"autoTimeUpdate;"`
	DeletedAt  gorm.DeletedAt
}

type PlayerRewardLog struct {
	PlayerRewardLogID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	PlayerId          uuid.UUID
	Player            Player `gorm:"ForeignKey:PlayerId"`
	RewardId          uuid.UUID
	Reward            Reward `gorm:"ForeignKey:RewardId"`
	CampaignId        uuid.UUID
	Campaign          Campaign  `gorm:"ForeignKey:CampaignId"`
	CreatedAt         time.Time `gorm:"autoTimeCreate;"`
	UpdatedAt         time.Time `gorm:"autoTimeUpdate;"`
	DeletedAt         gorm.DeletedAt
}

type PlayerActivityGameLog struct {
	PlayerActivityGameLogID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	PlayerId                uuid.UUID
	Player                  Player `gorm:"ForeignKey:PlayerId"`
	GameId                  uuid.UUID
	Game                    Game `gorm:"ForeignKey:GameId"`
	CampaignId              uuid.UUID
	Campaign                Campaign `gorm:"ForeignKey:CampaignId"`
	Point                   float64
	Stamina                 int64
	CreatedAt               time.Time `gorm:"autoTimeCreate;"`
	UpdatedAt               time.Time `gorm:"autoTimeUpdate;"`
	DeletedAt               gorm.DeletedAt
}
