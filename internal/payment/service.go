package payment

import (
	"context"
	"math/rand"
	"time"

	"github.com/danielalmeidafarias/go-saga/pkg"
	"github.com/google/uuid"
)

type PaymentService struct {
	validator         *pkg.Validator
	paymentRepository *PaymentRepository
}

func NewPaymentService(
	paymentRepository *PaymentRepository,
	validator *pkg.Validator,
) *PaymentService {
	return &PaymentService{
		validator:         validator,
		paymentRepository: paymentRepository,
	}
}

func (s *PaymentService) CreatePayment(ctx context.Context, params CreatePaymentParams) (*CreatePaymentResponse, error) {
	if err := s.validator.Validate(params); err != nil {
		return nil, err
	}

	input := CreatePaymentInput{
		Description: params.Description,
		Value:       params.Value,
	}

	payment, err := s.paymentRepository.Create(input)
	if err != nil {
		return nil, err
	}

	return &CreatePaymentResponse{
		PaymentUUID: payment.UUID.String(),
		Message:     "Payment created successfully",
	}, nil
}

func (s *PaymentService) GetPayment(ctx context.Context, id string) (*GetPaymentResponse, error) {
	paymentUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	payment, err := s.paymentRepository.GetOneByUUID(paymentUUID)
	if err != nil {
		return nil, err
	}

	return &GetPaymentResponse{
		UUID:        payment.UUID,
		Description: payment.Description,
		Value:       payment.Value,
		Status:      payment.Status,
	}, nil
}

// ProcessPayment simula um serviço de pagamento externo com comportamento assíncrono
// A função simula uma espera e pode falhar aleatoriamente
func (s *PaymentService) ProcessPayment(ctx context.Context, id string) (*ProcessPaymentResponse, error) {
	paymentUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	// Verifica se o pagamento existe
	_, err = s.paymentRepository.GetOneByUUID(paymentUUID)
	if err != nil {
		return nil, err
	}

	// Simula tempo de processamento assíncrono (1-5 segundos)
	processingTime := time.Duration(rand.Intn(5)+1) * time.Second
	time.Sleep(processingTime)

	// Simula falhas aleatórias (30% de chance de falhar)
	var newStatus PaymentStatus
	if rand.Float64() < 0.3 {
		newStatus = PaymentStatusFailed
	} else {
		newStatus = PaymentStatusProcessed
	}

	// Atualiza o status do pagamento
	updateInput := UpdatePaymentInput{
		Status: &newStatus,
	}

	updatedPayment, err := s.paymentRepository.Update(paymentUUID, updateInput)
	if err != nil {
		return nil, err
	}

	message := "Payment processed successfully"
	if newStatus == PaymentStatusFailed {
		message = "Payment processing failed"
	}

	return &ProcessPaymentResponse{
		PaymentUUID: updatedPayment.UUID.String(),
		Status:      updatedPayment.Status,
		Message:     message,
	}, nil
}
