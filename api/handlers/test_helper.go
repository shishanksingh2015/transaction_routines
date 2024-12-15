package handlers

import (
	"github.com/gofiber/fiber/v2"
	customError "routines/customerror"
)

func GetApp() *fiber.App {
	return fiber.New(fiber.Config{ErrorHandler: customError.CustomErrorHandler})
}
