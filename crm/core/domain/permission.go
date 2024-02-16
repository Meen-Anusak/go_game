package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type PermissionGroup struct {
	PermissionGroupID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();"`
	PermissionGroupName string
	Permission          []Permission
	CreatedAt           time.Time `gorm:"autoTimeCreate;"`
	UpdatedAt           time.Time `gorm:"autoTimeUpdate;"`
	DeletedAt           gorm.DeletedAt
}

type Permission struct {
	PermissionID      uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();"`
	PermissionName    string
	PermissionGroupId uuid.UUID
	PermissionGroup   PermissionGroup
	CreatedAt         time.Time `gorm:"autoTimeCreate;"`
	UpdatedAt         time.Time `gorm:"autoTimeUpdate;"`
	DeletedAt         gorm.DeletedAt
}
