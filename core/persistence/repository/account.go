package repository

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2/log"
	"routines/core/domain"
	"routines/core/persistence/mapper"
	"routines/customerror"
	"routines/db"
)

//go:generate mockgen -source=account.go  -destination=../../../mocks/core/repository/account.go -package=repository
type accountRepository struct {
	Db db.BaseDB
}

type AccountRepository interface {
	Create(ctx context.Context, request *domain.Account) error
}

func NewAccountRepository(dB *sql.DB) AccountRepository {
	return &accountRepository{Db: db.NewBaseDB(dB)}
}

// Create
//
//	@Description: Will create an entry in account table for given domain values
//	@param ctx context
//	@param request domain.Account which will be mapped to data access object
//	@return error
func (ar *accountRepository) Create(ctx context.Context, request *domain.Account) error {
	log.Info("mapping account domain to dao to save in database")
	dao := mapper.MapToAccountDao(*request)

	query := "insert into accounts(document_number) VALUES ($1)"

	err := ar.Db.Insert(ctx, query, dao.DocumentNumber)
	if err != nil {
		return customerror.InternalError(err.Error())
	}

	log.Info("successfully created account.")

	return nil
}
