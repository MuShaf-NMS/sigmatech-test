package helper

import "fmt"

type CustomError struct {
	Code   int
	Errors interface{}
}

func (ce *CustomError) Error() string {
	return fmt.Sprintf("%d; errors: %s", ce.Code, ce.Errors)
}

// Helper to create custom error
func NewError(statusCode int, errors interface{}) error {
	return &CustomError{
		Code:   statusCode,
		Errors: errors,
	}
}
