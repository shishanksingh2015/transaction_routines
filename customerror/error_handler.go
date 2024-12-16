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

// CustomErrorHandler
//
//	@Description: It is a custom error handler which will parse error depending on the type and it will work on middleware
//	@param ctx fiber context
//	@param err type of error RoutineError or fiber error
//	@return error
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

// BadRequest
//
//	@Description: Custom bad request with message
//	@param msg
//	@return RoutineError
func BadRequest(msg string) RoutineError {
	return newRoutineError(http.StatusBadRequest, msg)
}

// ConflictRequest
//
//	@Description: Custom conflict error with message
//	@param msg
//	@return RoutineError
func ConflictRequest(msg string) RoutineError {
	return newRoutineError(http.StatusConflict, msg)
}

// InternalError
//
//	@Description: Custom internal error with message
//	@param msg
//	@return RoutineError
func InternalError(msg string) RoutineError {
	return newRoutineError(http.StatusInternalServerError, msg)
}

// NotFound
//
//	@Description: Custom not found error with message
//	@param msg
//	@return RoutineError
func NotFound(msg string) RoutineError {
	return routineError{
		message:    msg,
		statusCode: http.StatusNotFound,
	}
}

// MethodNotAllowed
//
//	@Description: Custom method not found error with message
//	@param msg
//	@return RoutineError
func MethodNotAllowed(msg string) RoutineError {
	return routineError{
		message:    msg,
		statusCode: http.StatusMethodNotAllowed,
	}
}

func writeResponse(ctx *fiber.Ctx, err RoutineError) error {
	return ctx.Status(err.StatusCode()).JSON(err.Body())
}
