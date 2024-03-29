package domain

import (
	"github.com/google/uuid"
	"time"
)

type Player struct {
	PlayerID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"player_id"`
	PlayerName string    `json:"player_name" validate:"required,min=3"`
	CreatedAt  time.Time `gorm:"autoTimeCreate;" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoTimeUpdate;" json:"updated_at"`
}
