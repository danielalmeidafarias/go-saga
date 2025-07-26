package payment

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusProcessed PaymentStatus = "processed"
	PaymentStatusFailed    PaymentStatus = "failed"
)

type PaymentMethod string

const (
	PaymentMethodCreditCard    PaymentMethod = "CREDIT_CARD"
	PaymentMethodDigitalWallet PaymentMethod = "DIGITAL_WALLET"
)

type Payment struct {
	gorm.Model
	UUID          uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid()"`
	Description   string        `gorm:"not null;type:varchar(256);default:null"`
	Value         float64       `gorm:"not null;default:null"`
	Status        PaymentStatus `gorm:"type:varchar(20);default:'pending'"`
	PaymentMethod PaymentMethod `gorm:"type:varchar(20);default:null"`
}
