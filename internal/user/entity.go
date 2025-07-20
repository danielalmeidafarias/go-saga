package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name     string    `gorm:"not null; type:varchar(100); default:null"`
	Nickname *string   `gorm:"type:varchar(100); default:null"`
	Email    string    `gorm:"not null; unique; type:varchar(256); default:null"`
	Birthday time.Time `gorm:"not null"`
	Passport string    `gorm:"not null;unique;type:varchar(8);default:null"`
}
