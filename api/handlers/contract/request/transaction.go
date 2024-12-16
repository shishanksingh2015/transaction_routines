package request

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

// Transaction model info
type Transaction struct {
	// account id from user accounts
	AccountId     int     `json:"account_id" validate:"required"`
	OperationType int     `json:"operation_type" validate:"required"` // operation type ( 1,2,3,4)
	Amount        float64 `json:"amount" validate:"required"`         // amount in float with 2 decimal places 11.22
}

func (t *Transaction) Validate() error {
	err := validator.New().Struct(t)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
