package payment

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

type CreatePaymentInput struct {
	Description string
	Value       float64
}

type UpdatePaymentInput struct {
	Status *PaymentStatus `json:"status,omitempty"`
}

func NewPaymentRepository(
	db *gorm.DB,
) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (r *PaymentRepository) Create(in CreatePaymentInput) (*Payment, error) {
	payment := &Payment{
		Description: in.Description,
		Value:       in.Value,
		Status:      PaymentStatusPending,
	}

	err := r.db.Create(payment).Error
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (r *PaymentRepository) GetOneByUUID(uuid uuid.UUID) (*Payment, error) {
	var payment *Payment

	err := r.db.Where("uuid = ?", uuid).First(&payment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}

	return payment, nil
}

func (r *PaymentRepository) Update(uuid uuid.UUID, in UpdatePaymentInput) (*Payment, error) {
	payment, err := r.GetOneByUUID(uuid)
	if err != nil {
		return nil, err
	}

	if err := r.db.Model(payment).Updates(in).Error; err != nil {
		return nil, err
	}

	return payment, nil
}
