package game_model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Reward struct {
	RewardID    uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();"`
	RewardName  string
	Description string
	PointValue  float64
	Type        string
	CampaignId  uuid.UUID
	Campaign    Campaign `gorm:"ForeignKey:CampaignId"`
	BrandId     uuid.UUID
	Brand       Brand     `gorm:"ForeignKey:BrandId"`
	CreatedAt   time.Time `gorm:"autoTimeCreate;"`
	UpdatedAt   time.Time `gorm:"autoTimeUpdate;"`
	DeletedAt   gorm.DeletedAt
}
