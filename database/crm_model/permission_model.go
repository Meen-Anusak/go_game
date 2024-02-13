package crm_model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Permission struct {
	PermissionID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();"`
	PermissionName string
	CreatedAt      time.Time `gorm:"autoTimeCreate;"`
	UpdatedAt      time.Time `gorm:"autoTimeUpdate;"`
	DeletedAt      gorm.DeletedAt
}

type PermissionGroup struct {
	PermissionGroupID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();"`
	PermissionGroupName string
	Permission          []Permission
	CreatedAt           time.Time `gorm:"autoTimeCreate;"`
	UpdatedAt           time.Time `gorm:"autoTimeUpdate;"`
	DeletedAt           gorm.DeletedAt
}
