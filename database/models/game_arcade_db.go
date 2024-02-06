package models

type Items struct {
	Name  string `gorm:"name"`
	value uint   `gorm:"value"`
}
