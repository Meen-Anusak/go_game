package main

import (
	"go_game/database"
	"go_game/game/adapter"
	"go_game/game/core"
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

	// Player APP
	playerRepo := adapter.NewGormPlayerRepository(db)
	playerService := core.NewPlayerService(playerRepo)
	playerHandler := adapter.NewHttpPlayerHandler(playerService)
	// Player router

	app.Post("/player", playerHandler.CreateNewPlayerHandler)

	app.Listen(":8080")

}
