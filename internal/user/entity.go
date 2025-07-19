package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name     string    `gorm:"size=100"`
	Nickname *string   `gorm:"size=100"`
	Email    string    `gorm:"unique, size=256"`
	Birthday time.Time `gorm:"not null"`
	Passport string    `gorm:"unique"`
}
