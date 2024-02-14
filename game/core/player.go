package core

import (
	"github.com/google/uuid"
	"time"
)

type Player struct {
	PlayerID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"player_id"`
	PlayerName string    `json:"player_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
