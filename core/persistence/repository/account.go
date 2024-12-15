package repository

import (
	"context"
	"database/sql"
	"routines/api/handlers/contract/request"
	"routines/customerror"
	"routines/db"
)

//go:generate mockgen -source=account.go  -destination=../../../mocks/core/repository/account.go -package=repository
type accountRepository struct {
	Db db.BaseDB
}

type AccountRepository interface {
	Create(ctx context.Context, request *request.AccountRequest) error
}

func NewAccountRepository(dB *sql.DB) AccountRepository {
	return &accountRepository{Db: db.NewBaseDB(dB)}
}

func (ar *accountRepository) Create(ctx context.Context, request *request.AccountRequest) error {
	query := "insert into accounts(document_number) VALUES ($1)"
	err := ar.Db.Insert(ctx, query, request.DocumentNumber)
	if err != nil {
		return customerror.InternalError(err.Error())
	}

	return nil
}
