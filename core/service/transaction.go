package service

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
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
	log.Info("initiating create transaction and creating unique txn id for new transaction")
	id, err := data.CreateTransactionId()
	if err != nil {
		return customerror.InternalError(customerror.SomethingWentWrong)
	}

	log.Info("creating new domain transaction")
	transaction := &domain.Transaction{
		TransactionID: id,
	}

	log.Info("Add operation type in transaction")
	err = transaction.AddOperationType(data.OperationType(request.OperationType))
	if err != nil {
		return err
	}

	log.Info("checking if account exist in account table for account id")
	account, err := ts.accountProvider.GetAccountById(ctx, request.AccountId)
	if err != nil {
		return err
	}

	if account.Id == 0 {
		return customerror.BadRequest(fmt.Sprintf(customerror.AccountNotFound, request.AccountId))
	}

	log.Info("add account id in transaction")
	transaction.AddAccountId(account.Id)

	log.Info("add amount in transaction")
	err = transaction.AddAmount(request.Amount)
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("creating transaction with id %s", transaction.TransactionID.String()))
	err = ts.transactionRepo.CreateTransaction(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}
