package router

import (
	"github.com/gofiber/fiber/v2"
	"go_game/middleware"

	"go_game/crm/handler"

	"go_game/crm/core/services"

	"gorm.io/gorm"
)

func SetupCrmRouter(app *fiber.App, db *gorm.DB) {
	crmRep := handler.NewGormUserRepository(db)
	crmService := services.NewUserService(crmRep)
	crmHandler := handler.NewHttpUserHandler(crmService)

	crmRouter := app.Group("/user")

	crmRouter.Get("/list", middleware.AuthMiddleWare(), crmHandler.GetAllUser)
	crmRouter.Post("/", middleware.AuthMiddleWare(), crmHandler.CreateNewUser)

}
