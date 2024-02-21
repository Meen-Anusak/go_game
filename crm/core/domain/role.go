package domain

import (
	"github.com/google/uuid"
	"time"
)

type Role struct {
	RoleID            uuid.UUID        `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"role_id"`
	RoleName          string           `json:"role_name"`
	PermissionGroupId *uuid.UUID       `json:"permission_group_id"`
	PermissionGroup   *PermissionGroup `gorm:"ForeignKey:PermissionGroupId" json:"permission_group"`
	CreatedAt         time.Time        `gorm:"autoTimeCreate;" json:"created_at"`
	UpdatedAt         time.Time        `gorm:"autoTimeUpdate;" json:"updated_at"`
}
