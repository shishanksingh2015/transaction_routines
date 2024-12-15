package customerror

import (
	"fmt"
)

type routineError struct {
	error
	message    string
	statusCode int
}

type RoutineError interface {
	Body() map[string]interface{}
	StatusCode() int
	Error() string
}

func (r routineError) Body() map[string]any {
	return map[string]any{
		"errorCode":    r.statusCode,
		"errorMessage": r.message,
	}
}

func (r routineError) StatusCode() int {
	return r.statusCode
}

func (r routineError) Error() string {
	if r.error == nil {
		return r.message
	}
	return fmt.Sprintf("%s: %s", r.message, r.error.Error())
}
