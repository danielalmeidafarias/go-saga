package domain

import (
	"context"

	"github.com/google/uuid"
)

type PaymentStatus int

const (
	PROCESSING PaymentStatus = iota + 1
	CANCELED
	FINISHED
)

type Payment struct {
	Price     float64
	BookingID uuid.UUID
	Status    PaymentStatus
}

type CreatePaymentInput struct {
	BookingID uuid.UUID
	Price     uuid.UUID
}

type PaymentRepository interface {
	CreatePayment(ctx context.Context, in CreatePaymentInput) (paymentID string, err *Error)
	GetPayment(ctx context.Context, paymentID uuid.UUID) (Payment, *Error)
	UpdatePaymentStatus(ctx context.Context, paymentID uuid.UUID, status PaymentStatus) *Error
}
