package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/gofiber/swagger" //
	"routines/api/handlers"
)

func BindRoutes(app *fiber.App, handlers *handlers.Handlers) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("OK")
	})
	app.Get("v1/swagger/*", swagger.HandlerDefault)
	app.Post("v1/account", handlers.AccountHandler.CreateAccount)
	app.Get("v1/account/:accountId", handlers.AccountHandler.GetAccount)
	app.Post("v1/transaction", handlers.TransactionHandler.CreateTransaction)
}
