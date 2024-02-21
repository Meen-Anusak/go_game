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

	player := app.Group("/player")
	player.Get("list", middleware.AuthMiddleWare(), playerHandler.GetAllPlayer)
	player.Get(":id", middleware.AuthMiddleWare(), playerHandler.GetPlayerById)
	player.Post("", middleware.AuthMiddleWare(), playerHandler.CreateNewPlayer)
	player.Patch("", middleware.AuthMiddleWare(), playerHandler.UpdatePlayer)
	player.Delete("", middleware.AuthMiddleWare(), playerHandler.DeletePlayer)

}
