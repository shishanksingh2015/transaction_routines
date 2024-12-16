package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"routines/core/data"
	"routines/customerror"
	"testing"
)

func TestTransaction_AddAmount(t *testing.T) {
	id, err := data.CreateTransactionId()
	assert.NoError(t, err)

	transaction := &Transaction{
		TransactionID: id,
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
func TestTransaction_AddOperationType(t *testing.T) {
	id, err := data.CreateTransactionId()
	assert.NoError(t, err)

	transaction := &Transaction{
		TransactionID: id,
	}

	tests := []struct {
		name          string
		operationType data.OperationType
		error         error
		expectedError bool
	}{
		{
			name:          "Successfully Add PurchaseWithInstallments operation type to Transaction",
			operationType: data.PurchaseWithInstallments,
			error:         nil,
			expectedError: false,
		},
		{
			name:          "Successfully Add CreditVoucher operation type to Transaction ",
			operationType: data.CreditVoucher,
			error:         nil,
			expectedError: false,
		},
		{
			name:          "Fail to add invalid operation type",
			operationType: data.OperationType(-1),
			expectedError: true,
			error:         customerror.BadRequest(fmt.Sprintf(customerror.OperationNotValid, data.OperationType(-1).Int())),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := transaction.AddOperationType(tt.operationType)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
