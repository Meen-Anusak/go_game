package database

import uuid "github.com/satori/go.uuid"

type Brand struct {
	BrandId   uuid.UUID `gorm:"type:uuid;primaryKey;"`
	BrandName string
}
