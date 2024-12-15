package cmd

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/context"
	"routines/api/handlers"
	"routines/api/router"
	"routines/core/persistence/provider"
	"routines/core/persistence/repository"
	"routines/core/service"
	"routines/db"
)

func StartService(ctx context.Context, app *fiber.App) {
	handler := getHandler()
	router.BindRoutes(app, handler)
}

func getHandler() *handlers.Handlers {
	config, err := db.LoadConfig(".")
	if err != nil {
		return nil
	}

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
