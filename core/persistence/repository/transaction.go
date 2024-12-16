package repository

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2/log"
	"routines/core/domain"
	"routines/core/persistence/mapper"
	"routines/db"
)

//go:generate mockgen -source=transaction.go  -destination=../../../mocks/core/repository/transaction.go -package=repository
type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction *domain.Transaction) error
}

type transactionRepository struct {
	db.BaseDB
}

func NewTransactionRepository(sqlDb *sql.DB) TransactionRepository {
	return transactionRepository{db.NewBaseDB(sqlDb)}
}

func (t transactionRepository) CreateTransaction(ctx context.Context, transaction *domain.Transaction) error {
	log.Info("mapping transaction domain to dao to save in database")
	transactionDao := mapper.MapToTransactionDao(*transaction)

	query := "insert into transactions(id,account_id,operation_type,amount,event_date) VALUES ($1,$2,$3,$4,$5)"
	err := t.Insert(
		ctx,
		query,
		transactionDao.Id,
		transactionDao.AccountId,
		transactionDao.OperationType,
		transactionDao.Amount,
		transactionDao.EventDate)
	if err != nil {
		return err
	}

	log.Info("successfully created the transaction record")
	return nil
}
