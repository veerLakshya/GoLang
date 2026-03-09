package tictactoe

import "fmt"

type InvalidMoveError struct {
	Message string
}

func (e *InvalidMoveError) Error() string {
	return e.Message
}

func NewInvalidMoveError(format string, args ...interface{}) *InvalidMoveError {
	return &InvalidMoveError{
		Message: fmt.Sprintf(format, args...),
	}
}

//  can verify error type as errors.As()
