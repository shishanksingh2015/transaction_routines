package dao

import (
	"github.com/google/uuid"
	"time"
)

type TransactionDao struct {
	Id            uuid.UUID `json:"id"`
	AccountId     int       `json:"account_id"`
	OperationType int       `json:"operation_type"`
	Amount        float64   `json:"amount"`
	EventDate     time.Time `json:"event_date"`
}
