package crm_model

import "github.com/google/uuid"

type User struct {
	USerID      uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();"`
	UserName    string
	PhoneNumber string
}
