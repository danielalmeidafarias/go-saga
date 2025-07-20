package flight

import (
	"context"

	"github.com/danielalmeidafarias/go-saga/pkg"
	"github.com/google/uuid"
)

type FlightService struct {
	validator        *pkg.Validator
	flightRepository *FlightRepository
}

func NewFlightService(
	flightRepository *FlightRepository,
	validator *pkg.Validator,
) *FlightService {
	return &FlightService{
		validator:        validator,
		flightRepository: flightRepository,
	}
}

func (s *FlightService) CreateFlight(ctx context.Context, params CreateFlightParams) (*CreateFlightResponse, error) {
	if err := s.validator.Validate(params); err != nil {
		return nil, err
	}

	input := CreateFlightInput{
		Origin:     params.Origin,
		Destiny:    params.Destiny,
		FlightDate: params.FlightDate,
	}

	flight, err := s.flightRepository.Create(input)
	if err != nil {
		return nil, err
	}

	return &CreateFlightResponse{
		FlightUUID: flight.UUID.String(),
		Message:    "Flight created successfully",
	}, nil
}

func (s *FlightService) GetFlight(ctx context.Context, id string) (*GetFlightResponse, error) {
	flightUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	flight, err := s.flightRepository.GetOneByUUID(flightUUID)
	if err != nil {
		return nil, err
	}

	return &GetFlightResponse{
		UUID:       flight.UUID,
		Origin:     flight.Origin,
		Destiny:    flight.Destiny,
		FlightDate: flight.FlightDate,
		Seats:      flight.Seats,
	}, nil
}

func (s *FlightService) UpdateFlight(ctx context.Context, id string, params UpdateFlightParams) (*Flight, error) {
	flightUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	input := UpdateFlightInput{
		Origin:     params.Origin,
		Destiny:    params.Destiny,
		FlightDate: params.FlightDate,
	}

	flight, err := s.flightRepository.Update(flightUUID, input)
	if err != nil {
		return nil, err
	}

	return flight, nil
}

func (s *FlightService) DeleteFlight(ctx context.Context, id string) error {
	flightUUID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return s.flightRepository.Delete(flightUUID)
}

func (s *FlightService) CreateFlightSeat(ctx context.Context, flightID string, price float64) (*Seat, error) {
	flightUUID, err := uuid.Parse(flightID)
	if err != nil {
		return nil, err
	}

	seat, err := s.flightRepository.CreateSeat(flightUUID, price)
	if err != nil {
		return nil, err
	}

	return seat, nil
}

func (s *FlightService) GetFlightSeat(ctx context.Context, id string) (*Seat, error) {
	seatUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	seat, err := s.flightRepository.GetOneSeatByUUID(seatUUID)
	if err != nil {
		return nil, err
	}

	return seat, nil
}

func (s *FlightService) UpdateFlightSeat(ctx context.Context, id string, params UpdateSeatInput) (*Seat, error) {
	seatUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	seat, err := s.flightRepository.UpdateSeat(seatUUID, params)
	if err != nil {
		return nil, err
	}

	return seat, nil
}
