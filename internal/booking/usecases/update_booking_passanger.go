package usecases

/*
	REGRAS DE NEGÓCIO - UPDATE BOOKING PASSENGER

	1. VALIDAÇÃO DE ACESSO:
	   - Verificar se a reserva existe
	   - Verificar se o usuário tem permissão para alterar a reserva
	   - Apenas reservas "PENDING" ou "CONFIRMED" podem ser alteradas

	2. VALIDAÇÃO DOS DADOS:
	   - Verificar se o novo usuário existe no sistema
	   - Validar documentos e dados pessoais do novo passageiro
	   - Verificar se não há conflito com outras reservas

	3. REGRAS DE ALTERAÇÃO:
	   - Alteração só permitida até X horas antes do voo
	   - Pode haver taxa de alteração a ser cobrada
	   - Verificar política da companhia aérea para mudanças

	4. PROCESSO DE ATUALIZAÇÃO:
	   - Atualizar dados do passageiro na reserva
	   - Manter histórico da alteração para auditoria
	   - Recalcular preços se aplicável
	   - Gerar nova confirmação com dados atualizados

	5. NOTIFICAÇÕES:
	   - Notificar o novo passageiro sobre a reserva
	   - Notificar o passageiro original sobre a alteração
	   - Enviar nova confirmação por email
	   - Atualizar dados no sistema de check-in

	6. TRATAMENTO DE ERROS:
	   - Se novo usuário inválido: manter dados originais
	   - Se fora do prazo: bloquear alteração
	   - Rollback em caso de falha parcial
*/
