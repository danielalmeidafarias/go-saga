package usecases

import (
	"context"

	"github.com/danielalmeidafarias/go-saga/internal/booking/domain"
)

/*
	REGRAS DE NEGÓCIO - CREATE BOOKING

	1. VALIDAÇÃO INICIAL:
	   - Verificar se o assento existe e está disponível
	   - Verificar se o usuário existe e é válido
	   - Verificar se os dados de pagamento são válidos

	2. VERIFICAÇÃO DE DISPONIBILIDADE:
	   - O assento não deve estar ocupado por outra reserva ativa
	   - Verificar se o voo ainda aceita reservas (não partiu)
	   - Verificar capacidade máxima do voo

	3. PROCESSO DE RESERVA:
	   - Criar a reserva com status "PENDING"
	   - Reservar temporariamente o assento
	   - Iniciar processo de pagamento
	   - Gerar ID único para a reserva

	4. TRATAMENTO DE ERROS:
	   - Se assento indisponível: retornar erro específico
	   - Se usuário inválido: retornar erro de validação
	   - Se pagamento falha: cancelar reserva e liberar assento
	   - Implementar rollback em caso de falha

	5. SAGA PATTERN:
	   - Coordenar transações entre serviços (User, Flight, Payment)
	   - Implementar compensação em caso de falha
	   - Garantir consistência eventual entre serviços
*/

func (u *BookingUseCases) CreateBooking(ctx context.Context, p domain.CreateBookingParams) (*domain.Response, *domain.Error) {

	return nil, nil
}
