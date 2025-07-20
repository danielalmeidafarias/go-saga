package flight

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FlightRepository struct {
	db *gorm.DB
}

type CreateFlightInput struct {
	Destiny    string
	Origin     string
	FlightDate time.Time
}

type UpdateFlightInput struct {
	Destiny    *string    `json:"destiny,omitempty"`
	Origin     *string    `json:"origin,omitempty"`
	FlightDate *time.Time `json:"flightDate,omitempty"`
}

type UpdateSeatInput struct {
	Price    *float64 `json:"price,omitempty"`
	Reserved *bool    `json:"reserved,omitempty"`
}

func NewFlightRepository(
	db *gorm.DB,
) *FlightRepository {
	return &FlightRepository{
		db: db,
	}
}

func (r *FlightRepository) Create(in CreateFlightInput) (*Flight, error) {
	flight := &Flight{
		Origin:     in.Origin,
		Destiny:    in.Destiny,
		FlightDate: in.FlightDate,
	}

	err := r.db.Create(flight).Error
	if err != nil {
		return nil, err
	}

	return flight, nil
}

func (r *FlightRepository) GetOneByUUID(uuid uuid.UUID) (*Flight, error) {
	var flight *Flight

	err := r.db.Where("uuid = ?", uuid).First(flight).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("flight not found")
		}
		return nil, err
	}

	return flight, nil
}

func (r *FlightRepository) Update(uuid uuid.UUID, in UpdateFlightInput) (*Flight, error) {
	flight, err := r.GetOneByUUID(uuid)
	if err != nil {
		return nil, err
	}

	if err := r.db.Model(flight).Updates(in).Error; err != nil {
		return nil, err
	}

	return flight, nil
}

func (r *FlightRepository) Delete(uuid uuid.UUID) error {
	flight, err := r.GetOneByUUID(uuid)
	if err != nil {
		return err
	}

	err = r.db.Delete(flight).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *FlightRepository) CreateSeat(flightUUID uuid.UUID, price float64) (*Seat, error) {
	flight, err := r.GetOneByUUID(flightUUID)
	if err != nil {
		return nil, err
	}

	seat := &Seat{
		Price:    price,
		FlightID: flight.ID,
	}

	err = r.db.Create(seat).Error
	if err != nil {
		return nil, err
	}

	return seat, nil
}

func (r *FlightRepository) GetOneSeatByUUID(uuid uuid.UUID) (*Seat, error) {
	var seat *Seat

	if err := r.db.Where("uuid = ?", uuid).First(seat).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("seat not found")
		}

		return nil, err
	}

	return seat, nil
}

func (r *FlightRepository) UpdateSeat(uuid uuid.UUID, in UpdateSeatInput) (*Seat, error) {
	seat, err := r.GetOneSeatByUUID(uuid)
	if err != nil {
		return nil, err
	}

	err = r.db.Model(seat).Updates(in).Error
	if err != nil {
		return nil, err
	}

	return seat, nil
}

func (r *FlightRepository) DeleteSeat(uuid uuid.UUID) error {
	seat, err := r.GetOneSeatByUUID(uuid)
	if err != nil {
		return err
	}

	err = r.db.Delete(seat).Error
	if err != nil {
		return err
	}

	return nil
}
