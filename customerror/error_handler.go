package customerror

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func newRoutineError(code int, message string) RoutineError {
	return routineError{
		message:    message,
		statusCode: code,
	}
}

func BadRequest(msg string) RoutineError {
	return newRoutineError(http.StatusBadRequest, msg)
}
func ConflictRequest(msg string) RoutineError {
	return newRoutineError(http.StatusConflict, msg)
}
func InternalError(msg string) RoutineError {
	return newRoutineError(http.StatusInternalServerError, msg)
}

func NotFound(msg string) RoutineError {
	return routineError{
		message:    msg,
		statusCode: http.StatusNotFound,
	}
}

func MethodNotAllowed(msg string) RoutineError {
	return routineError{
		message:    msg,
		statusCode: http.StatusMethodNotAllowed,
	}
}

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	var routineErr RoutineError

	var fiberErr *fiber.Error

	switch true {
	case errors.As(err, &routineErr):
		return writeResponse(ctx, routineErr)
	case errors.As(err, &fiberErr):
		switch fiberErr.Code {
		case fiber.StatusNotFound:
			routineErr = NotFound(fiberErr.Error())
		case fiber.StatusMethodNotAllowed:
			routineErr = MethodNotAllowed(fiberErr.Error())
		default:
			routineErr = InternalError(fiberErr.Error())
		}
		return writeResponse(ctx, routineErr)
	default:
		return writeResponse(ctx, InternalError(err.Error()))
	}
}

func writeResponse(ctx *fiber.Ctx, err RoutineError) error {
	return ctx.Status(err.StatusCode()).JSON(err.Body())
}
