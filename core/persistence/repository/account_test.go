package repository

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"routines/core/domain"
	"routines/core/persistence/testhelper"
	"testing"
)

func TestAccountRepository_Create(t *testing.T) {
	ctx := context.Background()
	db, container, err := testhelper.RunPostgresContainer(ctx, "./../../../db/migration")
	assert.NoError(t, err)
	defer func(container testcontainers.Container, ctx context.Context) {
		err := container.Terminate(ctx)
		if err != nil {
			panic(fmt.Sprintf("unable to stop containers %s", container.GetContainerID()))
		}
	}(container, ctx)
	defer db.Close()

	accountRepo := NewAccountRepository(db)
	t.Run("Successfully create account record", func(t *testing.T) {
		requestAccount := &domain.Account{DocumentNumber: "12345678890"}
		err := accountRepo.Create(ctx, requestAccount)
		assert.NoError(t, err)
	})

	t.Run("should fail to create record with same document number", func(t *testing.T) {
		requestAccount := &domain.Account{DocumentNumber: "123456788910"}
		err := accountRepo.Create(ctx, requestAccount)
		assert.NoError(t, err)
		requestAccount1 := &domain.Account{DocumentNumber: "123456788910"}
		err = accountRepo.Create(ctx, requestAccount1)
		assert.Error(t, err)
	})
}
