package domain

import (
	"fmt"
	"routines/core/data"
	"routines/customerror"
)

type Transaction struct {
	TransactionID data.TransactionId
	OperationType data.OperationType
	Amount        float64
	AccountId     int
}

func (t *Transaction) AddAmount(amount float64) error {
	if data.IsValidOperationType(t.OperationType) {
		if !t.OperationType.IsCreditVoucher() {
			t.Amount = -amount
		} else {
			t.Amount = amount
		}
	} else {
		return customerror.BadRequest(fmt.Sprintf(customerror.AccountNotFound, t.OperationType.Int()))
	}

	return nil
}

func (t *Transaction) AddAccountId(accountId int) {
	t.AccountId = accountId
}
