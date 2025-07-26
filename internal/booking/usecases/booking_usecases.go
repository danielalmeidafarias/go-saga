package usecases

import (
	"github.com/danielalmeidafarias/go-saga/internal/booking/domain"
)

/*
	ESTRUTURA DOS USE CASES - BOOKING

	Esta struct centraliza todos os casos de uso relacionados às reservas de voo.
	Implementa o padrão SAGA para coordenar transações distribuídas entre os
	diferentes serviços (User, Flight, Payment).

	RESPONSABILIDADES:
	- Orquestrar operações complexas que envolvem múltiplos serviços
	- Garantir consistência eventual entre os serviços
	- Implementar lógica de compensação em caso de falhas
	- Validar regras de negócio antes de acionar outros serviços
	- Gerenciar estados intermediários durante as transações

	REPOSITÓRIOS UTILIZADOS:
	- bookingRepository: Gerencia dados das reservas
	- paymentRepository: Coordena com serviço de pagamento
	- flightRepository: Acessa informações de voos e assentos
	- seatRepository: Gerencia disponibilidade de assentos

	PADRÃO SAGA:
	- Cada operação é uma etapa da saga
	- Falhas em qualquer etapa disparam compensações
	- Estado da saga é mantido para recuperação
	- Timeouts e retries são implementados
*/

// type BookingUseCases interface {
// 	CreateBooking(ctx context.Context, p CreateBookingParams) (*Response, *Error)
// 	UpdateBooking(ctx context.Context, p UpdateBookingParams) (*Response, *Error)
// 	CancelBooking(ctx context.Context, bookingID string) (*Response, *Error)
// 	ConfirmBooking(ctx context.Context, bookingID string) (*Response, *Error)
// 	GetBooking(ctx context.Context, bookingID string) (*Response, *Error)
// 	ListFlightBookings(ctx context.Context, bookingID string) (*Response, *Error)
// 	ListUserBookings(ctx context.Context, bookingID string) (*Response, *Error)
// 	Booking(ctx context.Context, bookingID string) (*Response, *Error)
// }

type BookingUseCases struct {
	bookingRepository domain.BookingRepository
	paymentRepository domain.PaymentRepository
	flightRepository  domain.FlightRepository
	seatRepotory      domain.SeatRepository
}

func NewBookingUseCases(
	bookingRepository domain.BookingRepository,
	paymentRepository domain.PaymentRepository,
	flightRepository domain.FlightRepository,
	seatRepotory domain.SeatRepository,
) *BookingUseCases {
	return &BookingUseCases{
		bookingRepository: bookingRepository,
		paymentRepository: paymentRepository,
		flightRepository:  flightRepository,
		seatRepotory:      seatRepotory,
	}
}
