package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"routines/cmd"
	customError "routines/customerror"
	"routines/db"

	_ "routines/docs"
)

func main() {
	app := fiber.New(fiber.Config{
		Immutable:    true,
		ErrorHandler: customError.CustomErrorHandler,
	})
	config, err := db.LoadConfig(".")
	if err != nil {
		log.Fatalf("unable to load config %v", err)
	}
	cmd.StartService(context.Background(), app)
	err = app.Listen(config.ServerAddress)
	if err != nil {
		log.Fatalf("unable to start server %v", err)
	}
}
