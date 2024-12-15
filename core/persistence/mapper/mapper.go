package mapper

import (
	"routines/core/domain"
	"routines/core/persistence/dao"
	"time"
)

func MapToTransactionDao(transaction domain.Transaction) dao.TransactionDao {
	return dao.TransactionDao{
		Id:            transaction.TransactionID.UUID(),
		AccountId:     transaction.AccountId,
		OperationType: transaction.OperationType.Int(),
		Amount:        transaction.Amount,
		EventDate:     time.Now().UTC(),
	}
}

func MapToAccountDao(acc domain.Account) dao.AccountDao {
	return dao.AccountDao{
		Id:             acc.Id,
		DocumentNumber: acc.DocumentNumber,
	}
}

func MapToAccount(acc dao.AccountDao) *domain.Account {
	return &domain.Account{
		Id:             acc.Id,
		DocumentNumber: acc.DocumentNumber,
	}
}
