package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	request "routines/api/handlers/contract/request"
	"routines/core/domain"
	"routines/customerror"
	"routines/mocks/core/persistence/provider"
	"routines/mocks/core/repository"
	"testing"
)

func TestTransactionService_CreateTransaction(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockTransactionRepo := repository.NewMockTransactionRepository(mockCtrl)
	mockAccountProvider := provider.NewMockAccountProvider(mockCtrl)
	service := NewTransactionService(mockTransactionRepo, mockAccountProvider)
	ctx := context.Background()

	t.Run("Successfully create transaction record for given request", func(t *testing.T) {
		requestTxn := &request.Transaction{
			AccountId:     1,
			OperationType: 1,
			Amount:        20,
		}
		mockTransactionRepo.EXPECT().CreateTransaction(ctx, gomock.Any()).
			Times(1).Return(nil)
		mockAccountProvider.EXPECT().GetAccountById(ctx, requestTxn.AccountId).Times(1).
			Return(&domain.Account{Id: requestTxn.AccountId}, nil)
		err := service.CreateTransaction(ctx, requestTxn)
		assert.NoError(t, err)
	})

	t.Run("fail to create transaction due to invalid operation type", func(t *testing.T) {
		requestTxn := &request.Transaction{
			AccountId:     1,
			OperationType: -1,
			Amount:        20,
		}

		err := service.CreateTransaction(ctx, requestTxn)
		assert.Error(t, err)
		assert.Errorf(t, err, "operation type not allowed -1")
	})

	t.Run("fail to create transaction due to invalid accountId", func(t *testing.T) {
		requestTxn := &request.Transaction{
			AccountId:     1,
			OperationType: 1,
			Amount:        20,
		}
		mockAccountProvider.EXPECT().GetAccountById(ctx, requestTxn.AccountId).Times(1).
			Return(nil, sql.ErrNoRows)
		err := service.CreateTransaction(ctx, requestTxn)
		assert.Error(t, err)
		assert.Errorf(t, err, fmt.Sprintf(customerror.AccountNotFound, requestTxn.AccountId))
	})
}
