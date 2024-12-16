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

func (t *Transaction) AddOperationType(operationType data.OperationType) error {
	if !data.IsValidOperationType(operationType) {
		return customerror.BadRequest(fmt.Sprintf(customerror.OperationNotValid, t.OperationType.Int()))
	}
	t.OperationType = operationType
	return nil
}

func (t *Transaction) AddAmount(amount float64) error {
	if t.OperationType.IsCreditVoucher() {
		t.Amount = amount
	} else if t.OperationType.IsPurchaseOrWithdraw() {
		t.Amount = -amount
	} else {
		return customerror.BadRequest(fmt.Sprintf(customerror.UnableToAddAmount, t.OperationType.Int()))
	}
	
	return nil
}

func (t *Transaction) AddAccountId(accountId int) {
	t.AccountId = accountId
}
