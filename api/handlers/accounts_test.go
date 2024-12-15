package handlers

import (
	"bytes"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http/httptest"
	"routines/api/handlers/contract/request"
	"routines/commons/utils"
	"routines/core/domain"
	customError "routines/customerror"
	"routines/mocks/core/service"
	"testing"
)

func getApp() *fiber.App {
	return fiber.New(fiber.Config{ErrorHandler: customError.CustomErrorHandler})
}
func TestAccountHandler_CreateAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAccountService := service.NewMockAccountService(mockCtrl)
	accountHandler := NewAccountHandler(mockAccountService)
	app := getApp()

	t.Run("Should return status 201 when account created successfully", func(t *testing.T) {
		accountReq := request.AccountRequest{DocumentNumber: "test@123"}
		mockAccountService.EXPECT().CreateAccount(gomock.Any(), &accountReq).Times(1).Return(nil)
		app.Post("/account", accountHandler.CreateAccount)
		json, err := utils.StructToJson(&accountReq)
		assert.NoError(t, err)
		testRequest := httptest.NewRequest("POST", "/account", bytes.NewBufferString(json))
		testRequest.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(testRequest, -1)
		assert.NoError(t, err)
		assert.Equal(t, 201, resp.StatusCode)
	})

	t.Run("Should return status 400 when request is not correct", func(t *testing.T) {
		accountReq := request.AccountRequest{DocumentNumber: "test@123"}
		app.Post("/account", accountHandler.CreateAccount)
		json, err := utils.StructToJson(&accountReq)
		assert.NoError(t, err)
		testRequest := httptest.NewRequest("POST", "/account", bytes.NewBufferString(json))
		resp, err := app.Test(testRequest, -1)
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("Should return return error if unable to creat account", func(t *testing.T) {
		accountReq := request.AccountRequest{DocumentNumber: "test@123"}
		mockAccountService.EXPECT().CreateAccount(gomock.Any(), &accountReq).Times(1).Return(errors.New("something went wrong"))
		app.Post("/account", accountHandler.CreateAccount)
		json, err := utils.StructToJson(&accountReq)
		assert.NoError(t, err)
		testRequest := httptest.NewRequest("POST", "/account", bytes.NewBufferString(json))
		testRequest.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(testRequest, -1)
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)
	})

}

func TestAccountHandler_GetAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAccountService := service.NewMockAccountService(mockCtrl)
	accountHandler := NewAccountHandler(mockAccountService)
	app := getApp()

	t.Run("Should return status 200 when account is fetched successfully", func(t *testing.T) {
		mockAccountService.EXPECT().GetAccountById(gomock.Any(), 2).Times(1).
			Return(&domain.Account{DocumentNumber: "test@123", Id: 2}, nil)
		app.Get("/account/:accountId", accountHandler.GetAccount)
		testRequest := httptest.NewRequest("GET", "/account/2", nil)
		resp, err := app.Test(testRequest, -1)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("Should return status 400 when request id is not there", func(t *testing.T) {
		app.Get("/account/:accountId", accountHandler.GetAccount)
		testRequest := httptest.NewRequest("GET", "/account/shi", nil)
		resp, err := app.Test(testRequest, -1)
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)
		errResp := map[string]any{}
		err = utils.ReadInto(resp.Body, &errResp)
		assert.NoError(t, err)
		assert.Equal(t, "please send correct accountId", errResp["errorMessage"])
	})

	t.Run("Should return status 500 when request id is not there", func(t *testing.T) {
		app.Get("/account/:accountId", accountHandler.GetAccount)
		mockAccountService.EXPECT().GetAccountById(gomock.Any(), 2).Times(1).
			Return(nil, customError.InternalError(customError.SomethingWentWrong))
		testRequest := httptest.NewRequest("GET", "/account/2", nil)
		resp, err := app.Test(testRequest, -1)
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)
		errResp := map[string]any{}
		err = utils.ReadInto(resp.Body, &errResp)
		assert.NoError(t, err)
		assert.Equal(t, customError.SomethingWentWrong, errResp["errorMessage"])
	})

}
