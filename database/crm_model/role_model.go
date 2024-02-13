package crm_model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Role struct {
	RoleID            uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();"`
	RoleName          string
	PermissionGroupID uuid.UUID
	PermissionGroup   PermissionGroup `gorm:"ForeignKey:PermissionGroupID"`
	CreatedAt         time.Time       `gorm:"autoTimeCreate;"`
	UpdatedAt         time.Time       `gorm:"autoTimeUpdate;"`
	DeletedAt         gorm.DeletedAt
}
