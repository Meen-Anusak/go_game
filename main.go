package main

import (
	"go_game/database"
	_ "go_game/docs"
	"go_game/router"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	"log"

	"github.com/gofiber/fiber/v2"
)

// @title Game API
// @description This is a sample server for a game API.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

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
	app.Get("/swagger/*", swagger.HandlerDefault)

	//app.Get("/token", middleware.AuthorizationMiddleware)

	router.SetupAuthRouter(app, db)
	router.SetupGameRouter(app, db)

	_ = app.Listen(":8080")

}
