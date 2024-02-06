package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByUID(tx *gorm.DB, uid string) (*User, error)
	GetUserByID(userID string) (*User, error)
}

type User struct {
	GameID     uint      `gorm:"autoIncrement"`
	UserID     uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4();"`
	UID        string    `gorm:"not null;size:255;"`
	PlayerName string    `gorm:"size:30"`
	Language   string
	Coin       int `gorm:"default:0;"`
	AqPoint    int `gorm:"default:0;"`

	TotalTicket    int
	TotalExp       int
	Level          int `gorm:"not null;default:1"`
	ExtraStamina   int
	DefaultStamina int
	RegisterAt     time.Time
	CreatedAt      time.Time `gorm:"autoTimeCreate"`
	LastLoginAt    time.Time
	LastActivityAt time.Time
	UpdatedAt      time.Time `gorm:"autoTimeUpdate"`
	DeletedAt      gorm.DeletedAt
}
type UserRewardLog struct {
	UrwLogID       uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4();"`
	UserID         uuid.UUID
	BankTranTypeID uuid.UUID
	ItemID         *string `gorm:"size:5"`
	Amount         int     `gorm:"not null;"`
	// Event          UserRewardLogEvent `gorm:"type:user_reward_log_event;"`
	CreatedAt time.Time `gorm:"autoTimeCreate;"`
}

type UserInventory struct {
	User      User      `gorm:"forignKey:UserID"`
	UserID    uuid.UUID `gorm:"primaryKey"`
	ItemID    string    `gorm:"primaryKey"`
	Amount    int       `gorm:"not null"`
	UpdatedAt time.Time `gorm:"autoTimeUpdate"`
}
