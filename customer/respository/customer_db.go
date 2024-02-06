package respository

import "gorm.io/gorm"

type Customer struct {
	*gorm.DB
	Name string `json:"name"`
}
