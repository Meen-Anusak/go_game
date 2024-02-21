package domain

import (
	"github.com/google/uuid"
	"time"
)

type PermissionGroup struct {
	PermissionGroupID   uuid.UUID     `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"permission_group_id" `
	PermissionGroupName string        `json:"permission_group_name"`
	Permission          *[]Permission `json:"permission"`
	CreatedAt           time.Time     `gorm:"autoTimeCreate;" json:"created_at"`
	UpdatedAt           time.Time     `gorm:"autoTimeUpdate" json:"updated_at"`
}

type Permission struct {
	PermissionID      uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"permission_id"`
	PermissionName    string     `json:"permission_name"`
	PermissionGroupId *uuid.UUID `json:"permission_group_id"`
	PermissionGroup   *PermissionGroup
	CreatedAt         time.Time `gorm:"autoTimeCreate;" json:"created_at"`
	UpdatedAt         time.Time `gorm:"autoTimeUpdate;" json:"updated_at"`
}
