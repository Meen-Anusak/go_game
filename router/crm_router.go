package router

import (
	"github.com/gofiber/fiber/v2"
	"go_game/middleware"

	"go_game/crm/handler"

	"go_game/crm/core/services"

	"gorm.io/gorm"
)

func SetupCrmRouter(app *fiber.App, db *gorm.DB) {
	userRep := handler.NewGormUserRepository(db)
	userService := services.NewUserService(userRep)
	userHandler := handler.NewHttpUserHandler(userService)

	brandRep := handler.NewGormBrandRepository(db)
	brandService := services.NewBrandService(brandRep)
	brandHandler := handler.NewHttpBrandHandler(brandService)

	userRouter := app.Group("/user")
	userRouter.Get("/list", middleware.AuthMiddleWare(), userHandler.GetAllUser)
	userRouter.Post("/", middleware.AuthMiddleWare(), userHandler.CreateNewUser)

	brandRouter := app.Group("/brand")
	brandRouter.Get("/list", brandHandler.GetAllBrand)
	brandRouter.Post("/", brandHandler.CreateNewBrand)

}
