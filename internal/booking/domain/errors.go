package domain

import "fmt"

type Error struct {
	Code StatusCode
	Err  string
}

func NewError(err error, code StatusCode) *Error {
	if code == OK || code == CREATED {
		code = INTERNAL
	}

	return &Error{
		Err:  err.Error(),
		Code: code,
	}
}

func NewInternalError() *Error {
	return &Error{
		Err:  "an internal error has occurred",
		Code: INTERNAL,
	}
}

func NewNotFoundError(entity string) *Error {
	return &Error{
		Err:  fmt.Sprintf("%s not found", entity),
		Code: NOT_FOUND,
	}
}
