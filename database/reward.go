package database

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Reward struct {
	RewardId    uuid.UUID `gorm:"type:uuid;primaryKey;"`
	RewardName  string
	Description string
	PointValue  float64
	Type        string
	Campaign    Campaign
	Brand       Brand
	*gorm.Model
}
