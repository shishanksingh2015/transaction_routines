package service

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"routines/api/handlers/contract/request"
	"routines/core/domain"
	"routines/customerror"
	"routines/mocks/core/persistence/provider"
	"routines/mocks/core/repository"
	"testing"
)

func TestAccountService_CreateAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAccountRepo := repository.NewMockAccountRepository(mockCtrl)
	mockAccountProvider := provider.NewMockAccountProvider(mockCtrl)
	service := NewAccountService(mockAccountRepo, mockAccountProvider)
	ctx := context.Background()
	t.Run("successfully create account", func(t *testing.T) {
		requestAccount := &request.AccountRequest{DocumentNumber: "123123"}
		mockAccountProvider.EXPECT().GetAccountByDocumentNumber(gomock.Any(), requestAccount.DocumentNumber).
			Times(1).Return(nil, sql.ErrNoRows)
		mockAccountRepo.EXPECT().Create(gomock.Any(), requestAccount).
			Times(1).Return(nil)
		err := service.CreateAccount(ctx, requestAccount)
		assert.NoError(t, err)
	})

	t.Run("return error when unable to create account because account already exists", func(t *testing.T) {
		requestAccount := &request.AccountRequest{DocumentNumber: "123123"}
		mockAccountProvider.EXPECT().GetAccountByDocumentNumber(gomock.Any(), requestAccount.DocumentNumber).
			Times(1).Return(&domain.Account{}, nil)
		err := service.CreateAccount(ctx, requestAccount)
		assert.Error(t, err)
	})

	t.Run("return internal error when not able to create account", func(t *testing.T) {
		requestAccount := &request.AccountRequest{DocumentNumber: "123123"}
		mockAccountProvider.EXPECT().GetAccountByDocumentNumber(gomock.Any(), requestAccount.DocumentNumber).
			Times(1).Return(nil, sql.ErrNoRows)
		mockAccountRepo.EXPECT().Create(gomock.Any(), requestAccount).
			Times(1).Return(customerror.InternalError(customerror.SomethingWentWrong))
		err := service.CreateAccount(ctx, requestAccount)
		assert.Error(t, err)
	})

}

func TestAccountService_GetAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAccountRepo := repository.NewMockAccountRepository(mockCtrl)
	mockAccountProvider := provider.NewMockAccountProvider(mockCtrl)
	service := NewAccountService(mockAccountRepo, mockAccountProvider)
	ctx := context.Background()
	t.Run("successfully Fetch account by Id", func(t *testing.T) {
		requestId := 1
		expectedData := &domain.Account{Id: 1, DocumentNumber: "123456"}
		mockAccountProvider.EXPECT().GetAccountById(gomock.Any(), requestId).
			Times(1).Return(expectedData, nil)
		result, err := service.GetAccountById(ctx, requestId)
		assert.NoError(t, err)
		assert.Equal(t, expectedData.DocumentNumber, result.DocumentNumber)
	})

	t.Run("fail fetch account when account do not exist", func(t *testing.T) {
		requestId := 1
		mockAccountProvider.EXPECT().GetAccountById(gomock.Any(), requestId).
			Times(1).Return(nil, sql.ErrNoRows)
		_, err := service.GetAccountById(ctx, requestId)
		assert.Error(t, err)
	})
}
