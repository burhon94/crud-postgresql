package apiErrors

import "fmt"

type myError struct {
	text string
	err error
}

func ApiError(text string, err error) *myError {
	return &myError{text: text, err: err}
}

func (receiver *myError) Error() string {
	return fmt.Sprintf("error: %v", receiver.err.Error())
}

func (receiver *myError) Unwrap() error {
	return receiver.err
}
