package handlers

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http/httptest"
	"routines/api/handlers/contract/request"
	"routines/commons/utils"
	"routines/customerror"
	service2 "routines/mocks/core/service"
	"testing"
)

func TestTransactionHandler_CreateTransaction(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockTransactionService := service2.NewMockTransactionService(mockCtrl)
	transactionHandler := NewTransactionHandler(mockTransactionService)

	app := GetApp()

	t.Run("Should return status 201 when transaction created successfully", func(t *testing.T) {
		transactionReq := request.Transaction{
			AccountId:     1,
			OperationType: 1,
			Amount:        1500,
		}
		mockTransactionService.EXPECT().CreateTransaction(gomock.Any(), &transactionReq).Times(1).Return(nil)
		app.Post("/v1/transaction", transactionHandler.CreateTransaction)
		json, err := utils.StructToJson(&transactionReq)
		assert.NoError(t, err)
		testRequest := httptest.NewRequest("POST", "/v1/transaction", bytes.NewBufferString(json))
		testRequest.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(testRequest, -1)
		assert.NoError(t, err)
		assert.Equal(t, 201, resp.StatusCode)
	})

	t.Run("Should return status 400 when transaction request not correct", func(t *testing.T) {
		transactionReq := request.Transaction{
			AccountId: 1,
			Amount:    1500,
		}
		app.Post("v1/transaction", transactionHandler.CreateTransaction)
		json, err := utils.StructToJson(&transactionReq)
		assert.NoError(t, err)
		testRequest := httptest.NewRequest("POST", "/v1/transaction", bytes.NewBufferString(json))
		testRequest.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(testRequest, -1)
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("Should return status 500 when transaction creation fails", func(t *testing.T) {
		transactionReq := request.Transaction{
			AccountId:     1,
			OperationType: 1,
			Amount:        1500,
		}
		mockTransactionService.EXPECT().CreateTransaction(gomock.Any(), &transactionReq).
			Times(1).Return(customerror.InternalError(customerror.SomethingWentWrong))
		app.Post("/v1/transaction", transactionHandler.CreateTransaction)
		json, err := utils.StructToJson(&transactionReq)
		assert.NoError(t, err)
		testRequest := httptest.NewRequest("POST", "/v1/transaction", bytes.NewBufferString(json))
		testRequest.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(testRequest, -1)
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)
	})
}
