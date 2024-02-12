package main

import (
	"go_game/database"
	"go_game/game/adapter"
	"go_game/game/core"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db, err := database.NewDatabase()
	if err != nil {
		panic("failed to connect database")
	}

	// Player APP
	playerRepo := adapter.NewGormPlayerRepository(db)
	playerService := core.NewPlayerService(playerRepo)
	playerHandler := adapter.NewHttpPlayerHandler(playerService)

	app.Post("/player", playerHandler.CreateNewPlayer)

	app.Listen(":8080")

}
