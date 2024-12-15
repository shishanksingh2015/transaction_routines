package data

import (
	"github.com/google/uuid"
	"routines/commons/utils"
)

func (v TransactionId) String() string {
	return uuid.UUID(v).String()
}

func (v TransactionId) UUID() uuid.UUID {
	return uuid.UUID(v)
}

func CreateTransactionId() (TransactionId, error) {
	id, err := utils.GenerateUUID()
	if err != nil {
		return [16]byte{}, err
	}
	return TransactionId(id), err
}
