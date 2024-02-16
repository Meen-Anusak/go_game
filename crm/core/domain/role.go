package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Role struct {
	RoleID            uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();"`
	RoleName          string
	PermissionGroupId uuid.UUID
	PermissionGroup   PermissionGroup `gorm:"ForeignKey:PermissionGroupId"`
	CreatedAt         time.Time       `gorm:"autoTimeCreate;"`
	UpdatedAt         time.Time       `gorm:"autoTimeUpdate;"`
	DeletedAt         gorm.DeletedAt
}
