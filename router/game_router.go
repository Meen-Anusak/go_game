package router

import (
	"github.com/gofiber/fiber/v2"
	core "go_game/game/core/services"
	"go_game/game/handler"
	"gorm.io/gorm"
)

func SetupGameRouter(app *fiber.App, db *gorm.DB) {

	playerRepo := handler.NewGormPlayerRepository(db)
	playerService := core.NewPlayerService(playerRepo)
	playerHandler := handler.NewHttpPlayerHandler(playerService)

	player := app.Group("/player")
	player.Get("list", playerHandler.GetAllPlayer)
	player.Get(":id", playerHandler.GetPlayerById)
	player.Post("", playerHandler.CreateNewPlayer)
	player.Patch("", playerHandler.UpdatePlayer)

}
