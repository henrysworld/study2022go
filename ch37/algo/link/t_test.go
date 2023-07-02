package link

import (
	"fmt"
	"testing"
)

import (
	"errors"
)

type CustomError struct {
	message string
	code    int
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error Code: %d, Message: %s", e.code, e.message)
}

func someFunction() error {
	return &CustomError{
		message: "an error occurred",
		code:    500,
	}
}

func wrapError() error {
	err := someFunction()
	if err != nil {
		return fmt.Errorf("failed to execute someFunction: %w", err)
	}
	return nil
}

func Test1(t *testing.T) {
	err := wrapError()
	if err != nil {
		var customErr *CustomError
		if errors.As(err, &customErr) {
			fmt.Printf("Custom error occurred: %v\n", customErr)
		} else {
			fmt.Printf("An unknown error occurred: %v\n", err)
		}
	}
}
