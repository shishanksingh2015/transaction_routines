package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"routines/api/handlers/contract/request"
	"routines/core/domain"
	"routines/core/persistence/provider"
	"routines/core/persistence/repository"
	"routines/customerror"
)

//go:generate mockgen -source=account.go  -destination=../../mocks/core/service/account.go -package=service
type AccountService interface {
	CreateAccount(ctx context.Context, request *request.AccountRequest) error
	GetAccountById(ctx context.Context, accountId int) (*domain.Account, error)
}

type accountService struct {
	accountRepo     repository.AccountRepository
	accountProvider provider.AccountProvider
}

func NewAccountService(repo repository.AccountRepository, provider provider.AccountProvider) AccountService {
	return &accountService{accountRepo: repo, accountProvider: provider}
}
func (ah *accountService) CreateAccount(ctx context.Context, request *request.AccountRequest) error {
	_, err := ah.accountProvider.GetAccountByDocumentNumber(ctx, request.DocumentNumber)
	if err == nil {
		return customerror.ConflictRequest(
			fmt.Sprintf("account with document number : %s  exists", request.DocumentNumber))
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return customerror.InternalError(err.Error())
	}

	err = ah.accountRepo.Create(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func (ah *accountService) GetAccountById(ctx context.Context, accountId int) (*domain.Account, error) {
	account, err := ah.accountProvider.GetAccountById(ctx, accountId)
	if err != nil {
		return nil, err
	}

	return account, nil
}
