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

func TestCreateAccount(t *testing.T) {
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
	accountRepo := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepo, accountProvider)

	hdlrs := handlers.Handlers{
		AccountHandler: handlers.NewAccountHandler(accountService),
	}

	t.Run("Should create a account successfully", func(t *testing.T) {
		requestBody := `{
			"document_number": "12345678901"
		}`

		app := integration.SetupTestApp(hdlrs)

		req := httptest.NewRequest("POST", "/v1/account", bytes.NewBufferString(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	})

	t.Run("Should Get a account successfully by Id", func(t *testing.T) {
		app := integration.SetupTestApp(hdlrs)

		err := accountRepo.Create(ctx, &domain.Account{DocumentNumber: "1234567890"})
		assert.NoError(t, err)
		result, err := accountProvider.GetAccountByDocumentNumber(ctx, "1234567890")
		assert.NoError(t, err)
		req := httptest.NewRequest("GET", "/v1/account/"+fmt.Sprintf(`%d`, +result.Id), nil)

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

}
