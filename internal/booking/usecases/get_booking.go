package usecases

/*
	REGRAS DE NEGÓCIO - GET BOOKING

	1. VALIDAÇÃO DE ACESSO:
	   - Verificar se a reserva existe
	   - Verificar se o usuário tem permissão para visualizar a reserva
	   - Usuário só pode ver suas próprias reservas (ou admin pode ver todas)

	2. DADOS RETORNADOS:
	   - Informações completas da reserva (ID, status, data)
	   - Dados do assento (número, classe, localização)
	   - Informações do voo (origem, destino, horários)
	   - Status do pagamento
	   - Dados do passageiro

	3. TRATAMENTO DE STATUS:
	   - Exibir status atual da reserva (PENDING, CONFIRMED, CANCELLED)
	   - Mostrar histórico de mudanças de status se aplicável
	   - Indicar se reserva está próxima do vencimento

	4. INFORMAÇÕES ADICIONAIS:
	   - Calcular tempo restante para o voo
	   - Mostrar opções disponíveis (cancelar, alterar)
	   - Exibir política de cancelamento aplicável
	   - Links para check-in se aplicável

	5. SEGURANÇA:
	   - Não expor dados sensíveis de pagamento
	   - Logs de acesso para auditoria
	   - Validar autorização antes de retornar dados
*/
