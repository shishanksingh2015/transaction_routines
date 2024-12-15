package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	handlers "routines/api/handlers/contract/request"
	"routines/core/service"
	"routines/customerror"
)

type AccountHandler struct {
	accountService service.AccountService
}

func NewAccountHandler(service service.AccountService) AccountHandler {
	return AccountHandler{accountService: service}
}
func (ah *AccountHandler) CreateAccount(c *fiber.Ctx) error {
	req := &handlers.AccountRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	err := req.Validate()
	if err != nil {
		return err
	}

	err = ah.accountService.CreateAccount(c.UserContext(), req)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).SendString("account created successfully")
}

func (ah *AccountHandler) GetAccount(c *fiber.Ctx) error {
	accountId, err := c.ParamsInt("accountId", -1)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	if accountId == -1 {
		return customerror.BadRequest("please send correct accountId")
	}

	account, err := ah.accountService.GetAccountById(c.UserContext(), accountId)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(account)
}
