package request

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type AccountRequest struct {
	DocumentNumber string `json:"document_number" validate:"required"`
}

func (ar *AccountRequest) Validate() error {
	err := validator.New().Struct(ar)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
