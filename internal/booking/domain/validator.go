package domain

type Validator interface {
	Validate(any) *Error
}
