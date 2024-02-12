package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func NewDatabase() (*gorm.DB, error) {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error read .env file")
	}

	var DbHost string = os.Getenv("DATABASE_HOST")
	var DbName string = os.Getenv("DATABASE_NAME")
	var DbUser string = os.Getenv("DATABASE_USERNAME")
	var DbPassword string = os.Getenv("DATABASE_PASSWORD")
	var DbPort string = os.Getenv("DATABASE_POST")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DbHost, DbUser, DbPassword, DbName, DbPort)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		//PreferSimpleProtocol: configs.DB_DRY_RUN,
	}), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info), // or logger.Silent if you don't want logs
		Logger: logger.Default.LogMode(logger.Silent), // or logger.Silent if you don't want logs
	})

	if err != nil {
		return nil, err // instead of panic, return the error
	}

	return db, nil
}