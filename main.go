package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"go_game/database"
	"go_game/game/core"
	"go_game/game/handler"

	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	db, err := database.NewDatabase()
	if err != nil {
		panic("failed to connect database")
	}

	e := database.RunMigrations(db)

	if e != nil {
		log.Fatal(e)
	}

	playerRepo := handler.NewGormPlayerRepository(db)
	playerService := core.NewPlayerService(playerRepo)
	playerHandler := handler.NewHttpPlayerHandler(playerService)

	app.Get("/player", playerHandler.GetAllPlayer)
	app.Post("/player", playerHandler.CreateNewPlayer)

	app.Listen(":8080")

}
