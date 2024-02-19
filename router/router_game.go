package router

import (
	"github.com/gofiber/fiber/v2"
	core "go_game/game/core/services"
	"go_game/game/handler"
	"go_game/middleware"
	"gorm.io/gorm"
)

func SetupGameRouter(app *fiber.App, db *gorm.DB) {

	playerRepo := handler.NewGormPlayerRepository(db)
	playerService := core.NewPlayerService(playerRepo)
	playerHandler := handler.NewHttpPlayerHandler(playerService)

	app.Get("/player", middleware.Protect(), playerHandler.GetAllPlayer)
	app.Post("/player", playerHandler.CreateNewPlayer)

}
