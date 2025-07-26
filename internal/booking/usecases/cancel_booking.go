package usecases

/*
	REGRAS DE NEGÓCIO - CANCEL BOOKING

	1. VALIDAÇÃO DA RESERVA:
	   - Verificar se a reserva existe
	   - Verificar se a reserva pertence ao usuário solicitante
	   - Verificar se a reserva ainda pode ser cancelada (política de cancelamento)

	2. VERIFICAÇÃO DE STATUS:
	   - Apenas reservas com status "PENDING" ou "CONFIRMED" podem ser canceladas
	   - Reservas já canceladas não podem ser canceladas novamente
	   - Reservas de voos que já partiram não podem ser canceladas

	3. PROCESSO DE CANCELAMENTO:
	   - Alterar status da reserva para "CANCELLED"
	   - Liberar o assento para outras reservas
	   - Processar reembolso se aplicável
	   - Notificar o usuário sobre o cancelamento

	4. POLÍTICA DE REEMBOLSO:
	   - Verificar tempo até o voo (24h, 48h regras)
	   - Calcular valor do reembolso baseado na política
	   - Processar estorno no método de pagamento original

	5. SAGA PATTERN:
	   - Coordenar cancelamento entre serviços
	   - Garantir que assento seja liberado
	   - Garantir que pagamento seja estornado
	   - Implementar compensação se algum serviço falhar
*/
