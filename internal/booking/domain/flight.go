package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Seat struct {
	FlightID string  `json:"flightID"`
	Reserved bool    `json:"reserved"`
	Price    float64 `json:"price"`
}

type Flight struct {
	Id         uuid.UUID
	Destiny    string
	FlightDate time.Time
	Seats      []Seat
}

type SeatRepository interface {
	FindOneByID(ctx context.Context, seatID uuid.UUID) (*Seat, *Error)
	SetSeatReserved(ctx context.Context, seatID uuid.UUID, reserved bool)
}

type FlightRepository interface {
	FindOneByID(ctx context.Context, flightID uuid.UUID) *Error
}
