package router

import (
	"github.com/gofiber/fiber/v2"
	"go_game/auth/core/service"
	"go_game/auth/handler"
	"gorm.io/gorm"
)

func SetupAuthRouter(app *fiber.App, db *gorm.DB) {
	authRepo := handler.NewGormAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewHttpAuthHandler(authService)

	authRouter := app.Group("/auth")

	authRouter.Post("/register", authHandler.Register)

}
