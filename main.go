package main

import (
	"go_game/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	db, err := database.NewDatabase()
	if err != nil {
		panic("failed to connect database")
	}

	e := database.RunMigrations(db)

	if e != nil {
		log.Fatal(e)
	}

	app.Listen(":8080")

}
