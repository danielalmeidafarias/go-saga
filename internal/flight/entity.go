package flight

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	UUID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Origin     string    `gorm:"not null; type:varchar(100); default:null"`
	Destiny    string    `gorm:"not null;type:varchar(100);default:null"`
	FlightDate time.Time `gorm:"not null"`
	Seats      []Seat
}

type Seat struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Price    float64   `gorm:"not null; default:null"`
	Reserved bool      `gorm:"default:false"`
	FlightID uint
}
