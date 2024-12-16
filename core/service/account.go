package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
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

// CreateAccount
//
//	@Description: It will handle the request from handler and apply logic on domain and forward to account repository
//	@param ctx user context
//	@param request AccountRequest from account handler
//	@return error
func (ah *accountService) CreateAccount(ctx context.Context, request *request.AccountRequest) error {
	log.Info("initiating create account request " +
		"and checking if account exist for same document number")

	_, err := ah.accountProvider.GetAccountByDocumentNumber(ctx, request.DocumentNumber)
	if err == nil {
		return customerror.ConflictRequest(
			fmt.Sprintf(customerror.AccountExists, request.DocumentNumber))
	}
	if !errors.Is(err, sql.ErrNoRows) {
		fmt.Errorf("database", err.Error())
		return customerror.InternalError(err.Error())
	}

	account := &domain.Account{
		DocumentNumber: request.DocumentNumber,
	}

	log.Info("checking for document number to be valid ")

	if !account.IsDocumentValid() {
		return customerror.BadRequest(customerror.DocumentNotValid)
	}

	log.Info("document number valid, creating account...")

	err = ah.accountRepo.Create(ctx, account)
	if err != nil {
		return err
	}

	return nil
}

// GetAccountById
//
//	@Description: Fetch account from provider for given account id
//	@param ctx
//	@param accountId
//	@return *domain.Account return the domain.Account from provider
//	@return error
func (ah *accountService) GetAccountById(ctx context.Context, accountId int) (*domain.Account, error) {
	log.Info(fmt.Sprintf("fetching account from provider for id %d exist", accountId))
	account, err := ah.accountProvider.GetAccountById(ctx, accountId)
	if err != nil {
		return nil, err
	}

	return account, nil
}
