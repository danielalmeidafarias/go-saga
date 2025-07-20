package payment

import (
	"github.com/google/uuid"
)

type CreatePaymentParams struct {
	Description string  `validate:"required"`
	Value       float64 `validate:"required,gt=0"`
}

type ProcessPaymentParams struct {
	Status PaymentStatus `validate:"required,oneof=processed failed"`
}

type CreatePaymentResponse struct {
	PaymentUUID string `json:"paymentUUID"`
	Message     string `json:"message"`
}

type GetPaymentResponse struct {
	UUID        uuid.UUID     `json:"paymentUUID"`
	Description string        `json:"description"`
	Value       float64       `json:"value"`
	Status      PaymentStatus `json:"status"`
}

type ProcessPaymentResponse struct {
	PaymentUUID string        `json:"paymentUUID"`
	Status      PaymentStatus `json:"status"`
	Message     string        `json:"message"`
}
