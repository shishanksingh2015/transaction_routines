package domain

import (
	"github.com/stretchr/testify/assert"
	"routines/core/data"
	"testing"
)

func TestTransaction_AddAmount(t *testing.T) {
	id, err := data.CreateTransactionId()
	assert.NoError(t, err)

	transaction := &Transaction{
		TransactionID: id,
		AccountId:     1,
	}

	tests := []struct {
		name           string
		operationType  data.OperationType
		amount         float64
		expectedAmount float64
		expectError    bool
	}{
		{
			name:           "Successfully Add amount to Transaction for negative operation",
			operationType:  data.PurchaseWithInstallments,
			amount:         10,
			expectedAmount: -10.00,
			expectError:    false,
		},
		{
			name:           "Successfully Add amount to Transaction for positive operation",
			operationType:  data.CreditVoucher,
			amount:         10,
			expectedAmount: 10.00,
			expectError:    false,
		},
		{
			name:          "Fail to add amount when operation type is invalid",
			operationType: data.OperationType(-1),
			amount:        10,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transaction.OperationType = tt.operationType

			err := transaction.AddAmount(tt.amount)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedAmount, transaction.Amount)
			}
		})
	}
}
