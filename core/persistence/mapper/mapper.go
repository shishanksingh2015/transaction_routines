package mapper

import (
	"routines/core/domain"
	"routines/core/persistence/dao"
	"time"
)

// MapToTransactionDao
//
//	@Description: It will map domain.Transaction to Data access object Transaction
//	@param transaction  domain object
//	@return dao.TransactionDao
func MapToTransactionDao(transaction domain.Transaction) dao.TransactionDao {
	return dao.TransactionDao{
		Id:            transaction.TransactionID.UUID(),
		AccountId:     transaction.AccountId,
		OperationType: transaction.OperationType.Int(),
		Amount:        transaction.Amount,
		EventDate:     time.Now().UTC(),
	}
}

// MapToAccountDao
//
//	@Description: It will map domain.Account to Data access object Account
//	@param Account  domain object
//	@return dao.AccountDao
func MapToAccountDao(acc domain.Account) dao.AccountDao {
	return dao.AccountDao{
		Id:             acc.Id,
		DocumentNumber: acc.DocumentNumber,
	}
}

// MapToAccount
//
//	@Description:  It will map Data access object Account to domain.Account
//	@param acc Data access object Account
//	@return *domain.Account
func MapToAccount(acc dao.AccountDao) *domain.Account {
	return &domain.Account{
		Id:             acc.Id,
		DocumentNumber: acc.DocumentNumber,
	}
}
