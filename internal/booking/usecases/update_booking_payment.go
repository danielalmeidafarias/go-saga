package usecases

/*
	REGRAS DE NEGÓCIO - UPDATE BOOKING PAYMENT

	1. VALIDAÇÃO DA RESERVA:
	   - Verificar se a reserva existe
	   - Verificar se o usuário tem permissão para alterar o pagamento
	   - Apenas reservas "PENDING" podem ter pagamento alterado

	2. VALIDAÇÃO DO NOVO PAGAMENTO:
	   - Verificar se o novo método de pagamento é válido
	   - Validar dados do cartão/conta bancária
	   - Verificar se o valor corresponde ao da reserva
	   - Confirmar disponibilidade de limite/saldo

	3. PROCESSO DE ALTERAÇÃO:
	   - Cancelar/estornar transação anterior se necessário
	   - Processar novo pagamento
	   - Atualizar dados de pagamento na reserva
	   - Manter histórico de transações

	4. TRATAMENTO DE FALHAS:
	   - Se novo pagamento falha: manter método anterior
	   - Implementar retry para falhas temporárias
	   - Notificar usuário sobre sucesso/falha
	   - Reverter alterações em caso de erro

	5. REGRAS FINANCEIRAS:
	   - Verificar se há diferença de valores a cobrar/estornar
	   - Aplicar taxas de alteração se aplicável
	   - Respeitar políticas de estorno do meio de pagamento
	   - Atualizar status financeiro da reserva

	6. SAGA PATTERN:
	   - Coordenar com serviço de pagamento
	   - Garantir consistência entre reserva e pagamento
	   - Implementar compensação para falhas
*/
