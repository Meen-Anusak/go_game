package database

import (
	"go_game/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitalDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=Game port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db

}
func RunMigrations(db *gorm.DB) error {
	err := db.Transaction(func(d *gorm.DB) error {
		err := d.AutoMigrate(&models.Items{})
		return err
	})

	return err
}
