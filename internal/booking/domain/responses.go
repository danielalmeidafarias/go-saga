package domain

type Response struct {
	Message string
	Code    StatusCode
}

func NewResponse(message string, code StatusCode) *Response {
	return &Response{
		Message: message,
		Code:    code,
	}
}
