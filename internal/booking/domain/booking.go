package domain

import (
	"context"

	"github.com/google/uuid"
)

type Booking struct {
	SeatID    uuid.UUID `json:"seat"`
	PaymentID uuid.UUID `json:"payment"`
	UserID    uuid.UUID `json:"userID"`
}

type CreateBookingInput struct {
	SeatID    *string
	PaymentID *string
	UserID    *string
}

type UpdateBookingInput struct {
	PaymentID *string
	UserID    *string
}

type BookingRepository interface {
	Create(ctx context.Context, in CreateBookingInput) (string, *Error)
	Get(ctx context.Context, id string) (*Booking, *Error)
	Update(ctx context.Context, in UpdateBookingInput) *Error
	Delete(ctx context.Context, id string) *Error
}
