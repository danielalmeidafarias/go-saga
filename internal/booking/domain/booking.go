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
	FindOne(ctx context.Context, id string) (*Booking, *Error)
	FindManyByUserID(ctx context.Context, userID string) ([]*Booking, *Error)
	Update(ctx context.Context, in UpdateBookingInput) *Error
	Delete(ctx context.Context, id string) *Error
}

type CreateBookingParams struct {
	SeatID    uuid.UUID
	UserID    uuid.UUID
	PaymentID *uuid.UUID
}

type UpdateBookingParams struct {
	UserID    *uuid.UUID
	PaymentID *uuid.UUID
}

type BookingUseCases interface {
	CreateBooking(ctx context.Context, p CreateBookingParams) (*Response, *Error)
	UpdateBooking(ctx context.Context, p UpdateBookingParams) (*Response, *Error)
	CancelBooking(ctx context.Context, bookingID string) (*Response, *Error)
	ConfirmBooking(ctx context.Context, bookingID string) (*Response, *Error)
	GetBooking(ctx context.Context, bookingID string) (*Response, *Error)
	ListFlightBookings(ctx context.Context, bookingID string) (*Response, *Error)
	ListUserBookings(ctx context.Context, bookingID string) (*Response, *Error)
	Booking(ctx context.Context, bookingID string) (*Response, *Error)
}
