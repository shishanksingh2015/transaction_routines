package request

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type Transaction struct {
	AccountId     int     `json:"account_id" validate:"required"`
	OperationType int     `json:"operation_type" validate:"required"`
	Amount        float64 `json:"amount" validate:"required"`
}

func (t *Transaction) Validate() error {
	err := validator.New().Struct(t)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
