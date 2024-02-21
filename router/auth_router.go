package router

import (
	"github.com/gofiber/fiber/v2"
	"go_game/auth/core/services"
	"go_game/auth/handler"
	"gorm.io/gorm"
)

func SetupAuthRouter(app *fiber.App, db *gorm.DB) {
	authRepo := handler.NewGormAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authHandler := handler.NewHttpAuthHandler(authService)

	authRouter := app.Group("/auth")

	authRouter.Post("/login", authHandler.Login)
	authRouter.Post("/register", authHandler.Register)

}
