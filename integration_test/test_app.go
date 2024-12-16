package integration

import (
	"github.com/gofiber/fiber/v2"
	"routines/api/handlers"
	customError "routines/customerror"
)

func SetupTestApp(handlers handlers.Handlers) *fiber.App {
	app := fiber.New(fiber.Config{ErrorHandler: customError.CustomErrorHandler})

	app.Post("/v1/account", handlers.AccountHandler.CreateAccount)
	app.Post("/v1/transaction", handlers.TransactionHandler.CreateTransaction)
	app.Get("/v1/account/:accountId", handlers.AccountHandler.GetAccount)
	return app
}
