package flight

import (
	"time"

	"github.com/google/uuid"
)

type CreateFlightParams struct {
	Origin     string    `validate:"required"`
	Destiny    string    `validate:"required"`
	FlightDate time.Time `validate:"required"`
}

type UpdateFlightParams struct {
	Origin     *string
	Destiny    *string
	FlightDate *time.Time
}

type CreateFlightResponse struct {
	FlightUUID string `json:"flightUUID"`
	Message    string `json:"message"`
}

type GetFlightResponse struct {
	UUID       uuid.UUID `json:"flightUUID"`
	Origin     string    `json:"origin"`
	Destiny    string    `json:"destiny"`
	FlightDate time.Time `json:"flightDate"`
	Seats      []Seat    `json:"seats"`
}
