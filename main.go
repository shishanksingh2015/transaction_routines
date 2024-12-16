package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"routines/cmd"
	customError "routines/customerror"
	"routines/db"

	_ "routines/docs"
)

func main() {
	app := New()
	config, err := db.LoadConfig()
	if err != nil {
		log.Fatalf("unable to load config %v", err)
	}

	log.Println("Server is starting")
	cmd.StartService(app, config)

	log.Println("Server is running at : " + config.ServerAddress)
	err = app.Listen(config.ServerAddress)
	if err != nil {
		log.Fatalf("unable to start server %v", err)
	}
}

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		Immutable:    true,
		ErrorHandler: customError.CustomErrorHandler,
	})

	app.Use(logger.New(cmd.LoggerConfig()))
	return app
}
