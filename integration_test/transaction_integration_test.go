package integration_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"net/http/httptest"
	"routines/api/handlers"
	"routines/core/domain"
	"routines/core/persistence/provider"
	"routines/core/persistence/repository"
	"routines/core/persistence/testhelper"
	"routines/core/service"
	integration "routines/integration_test"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	ctx := context.Background()
	db, container, err := testhelper.RunPostgresContainer(ctx, "./../db/migration")
	assert.NoError(t, err)
	defer func(container testcontainers.Container, ctx context.Context) {
		err := container.Terminate(ctx)
		if err != nil {
			panic(fmt.Sprintf("unable to stop containers %s", container.GetContainerID()))
		}
	}(container, ctx)
	defer db.Close()

	accountProvider := provider.NewAccountProvider(db)
	transactionRepo := repository.NewTransactionRepository(db)
	accountRepo := repository.NewAccountRepository(db)
	transactionService := service.NewTransactionService(transactionRepo, accountProvider)

	hdlrs := handlers.Handlers{
		TransactionHandler: handlers.NewTransactionHandler(transactionService),
	}

	t.Run("Should create a transaction successfully", func(t *testing.T) {
		err = accountRepo.Create(ctx, &domain.Account{DocumentNumber: "12345678890"})
		assert.NoError(t, err)
		resultAcc, err := accountProvider.GetAccountByDocumentNumber(ctx, "12345678890")
		assert.NoError(t, err)
		app := integration.SetupTestApp(hdlrs)

		requestBody := `{
            "account_id":` + fmt.Sprintf(`%d`, resultAcc.Id) + `,
			"operation_type": 1,
			"amount": 100.0
		}`
		req := httptest.NewRequest("POST", "/v1/transaction", bytes.NewBufferString(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	})

}
