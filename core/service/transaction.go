package service

import (
	"context"
	"fmt"
	"routines/api/handlers/contract/request"
	"routines/core/data"
	"routines/core/domain"
	"routines/core/persistence/provider"
	"routines/core/persistence/repository"
	"routines/customerror"
)

//go:generate mockgen -source=transaction.go  -destination=../../mocks/core/service/transaction.go -package=service
type TransactionService interface {
	CreateTransaction(ctx context.Context, request *request.Transaction) error
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
	accountProvider provider.AccountProvider
}

func NewTransactionService(repo repository.TransactionRepository, acProvider provider.AccountProvider) TransactionService {
	return &transactionService{
		transactionRepo: repo,
		accountProvider: acProvider,
	}
}
func (ts *transactionService) CreateTransaction(ctx context.Context, request *request.Transaction) error {
	id, err := data.CreateTransactionId()
	if err != nil {
		return customerror.InternalError(customerror.SomethingWentWrong)
	}

	transaction := &domain.Transaction{
		TransactionID: id,
		OperationType: data.OperationType(request.OperationType),
	}

	account, err := ts.accountProvider.GetAccountById(ctx, request.AccountId)
	if err != nil {
		return err
	}

	if account.Id == 0 {
		return customerror.BadRequest(fmt.Sprintf(customerror.AccountNotFound, request.AccountId))
	}

	transaction.AddAccountId(account.Id)
	err = transaction.AddAmount(request.Amount)
	if err != nil {
		return err
	}

	err = ts.transactionRepo.CreateTransaction(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}
