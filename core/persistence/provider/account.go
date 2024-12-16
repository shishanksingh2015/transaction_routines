package provider

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"routines/core/domain"
	"routines/core/persistence/dao"
	"routines/core/persistence/mapper"
	"routines/customerror"
	"routines/db"
)

//go:generate mockgen -source=account.go  -destination=../../../mocks/core/persistence/provider/account.go -package=provider
type AccountProvider interface {
	GetAccountById(ctx context.Context, data int) (*domain.Account, error)
	GetAccountByDocumentNumber(ctx context.Context, data interface{}) (*domain.Account, error)
}

type accountProvider struct {
	dB db.BaseDB
}

func NewAccountProvider(sqlDb *sql.DB) AccountProvider {
	return &accountProvider{dB: db.NewBaseDB(sqlDb)}
}

// GetAccountById
//
//	@Description: Retrieves an account for a user based on the provided account ID.
//	@param ctx context
//	@param data the unique identifier of the account (int).
//	@return *domain.Account-the account details mapped from the data access object (DAO).
//	@return error
func (a *accountProvider) GetAccountById(ctx context.Context, data int) (*domain.Account, error) {
	resultDao := &dao.AccountDao{}
	query := "SELECT * FROM accounts WHERE id=$1"
	err := a.dB.DB(ctx).QueryRowContext(ctx, query, data).
		Scan(&resultDao.Id, &resultDao.DocumentNumber, &resultDao.CreatedAt)

	if err == nil {
		log.Info(fmt.Sprintf("account with account id %d found", data))
		result := mapper.MapToAccount(*resultDao)

		return result, nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, customerror.BadRequest(fmt.Sprintf(customerror.AccountNotFound, data))
	}

	return nil, customerror.InternalError(customerror.SomethingWentWrong)
}

// GetAccountByDocumentNumber
//
//	@Description: Retrieves an account for a user based on the provided document number.
//	@param ctx context
//	@param data the unique identifier of the account (document number).
//	@return *domain.Account-the account details mapped from the data access object (DAO).
//	@return error
func (a *accountProvider) GetAccountByDocumentNumber(ctx context.Context, data interface{}) (*domain.Account, error) {
	resultDao := &dao.AccountDao{}

	query := "SELECT * FROM accounts WHERE document_number=$1"

	err := a.dB.DB(ctx).QueryRowContext(ctx, query, data).
		Scan(&resultDao.Id, &resultDao.DocumentNumber, &resultDao.CreatedAt)

	if err != nil {
		return nil, err
	}

	log.Info(fmt.Sprintf("account with account document number %s found", data))
	result := mapper.MapToAccount(*resultDao)

	return result, nil
}
