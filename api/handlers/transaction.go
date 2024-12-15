package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	handlers "routines/api/handlers/contract/request"
	"routines/core/service"
	"routines/customerror"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) TransactionHandler {
	return TransactionHandler{transactionService: transactionService}
}

// CreateTransaction godoc
//
// @Summary 	Initiate Transaction for a payment
// @Description	Create a transaction for given operation type with amount and account id
// @Accept	 	application/json
// @Param	  	data					body	handlers.Transaction	true "Transaction"
// @Produce		json
// @Success		201
// @Failure		400 {object}	interface{}
// @Failure		500 {object}	interface{}
// @Router      /transaction [post]
func (th *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	req := &handlers.Transaction{}
	if err := c.BodyParser(req); err != nil {
		return err
	}

	err := req.Validate()
	if err != nil {
		return customerror.BadRequest(err.Error())
	}

	err = th.transactionService.CreateTransaction(c.UserContext(), req)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).SendString("transaction created successfully")
}
