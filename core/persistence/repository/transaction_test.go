package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"routines/core/data"
	"routines/core/domain"
	"routines/core/persistence/provider"
	"routines/core/persistence/testhelper"
	"testing"
)

func TestTransactionRepository_CreateTransaction(t *testing.T) {
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

	transactionRepo := NewTransactionRepository(db)
	accountRepo := NewAccountRepository(db)
	accountProvider := provider.NewAccountProvider(db)
	err = accountRepo.Create(ctx, &domain.Account{DocumentNumber: "12345678890"})
	assert.NoError(t, err)
	resultAcc, err := accountProvider.GetAccountByDocumentNumber(ctx, "12345678890")
	assert.NoError(t, err)

	t.Run("should successfully save transaction", func(t *testing.T) {
		txnDomain := &domain.Transaction{
			TransactionID: data.TransactionId(uuid.New()),
			OperationType: data.Withdrawal,
			Amount:        1500,
			AccountId:     resultAcc.Id,
		}
		err := transactionRepo.CreateTransaction(ctx, txnDomain)
		assert.NoError(t, err)
	})

	t.Run("should not create transaction if account number do not exists", func(t *testing.T) {
		txnDomain := &domain.Transaction{
			TransactionID: data.TransactionId(uuid.New()),
			OperationType: data.Withdrawal,
			Amount:        1500,
			AccountId:     11,
		}
		err := transactionRepo.CreateTransaction(ctx, txnDomain)
		assert.Error(t, err)
	})
}
