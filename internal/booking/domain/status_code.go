package domain

type StatusCode int

const (
	CREATED StatusCode = iota + 1
	OK
	NOT_FOUND
	BAD_REQUEST
	INTERNAL
	FAILED_PRECONDITION
)
