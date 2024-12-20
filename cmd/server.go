package cmd

import (
	"github.com/gofiber/fiber/v2"
	"routines/api/handlers"
	"routines/api/router"
	"routines/core/persistence/provider"
	"routines/core/persistence/repository"
	"routines/core/service"
	"routines/db"
)

func StartService(app *fiber.App, config db.Config) {
	handler := getHandler(config)
	router.BindRoutes(app, handler)
}

// getHandler
//
//	@Description: It is a private method that will generate all handlers and services and other dependencies
//	@param config Db config to connect to db
//	@return *handlers.Handlers list of handlers
func getHandler(config db.Config) *handlers.Handlers {
	database := db.ConnectDb(config)

	accountRepo := repository.NewAccountRepository(database)
	accountProvider := provider.NewAccountProvider(database)
	accountService := service.NewAccountService(accountRepo, accountProvider)

	txnRepo := repository.NewTransactionRepository(database)
	txnService := service.NewTransactionService(txnRepo, accountProvider)
	txnHandler := handlers.NewTransactionHandler(txnService)

	return &handlers.Handlers{
		AccountHandler:     handlers.NewAccountHandler(accountService),
		TransactionHandler: txnHandler,
	}
}
