package usecases

/*
	REGRAS DE NEGÓCIO - LIST BOOKINGS BY FLIGHT

	1. VALIDAÇÃO DE ACESSO:
	   - Verificar se o voo existe
	   - Apenas funcionários/admin podem listar todas as reservas de um voo
	   - Usuários comuns só veem suas próprias reservas

	2. FILTROS E PAGINAÇÃO:
	   - Permitir filtro por status (PENDING, CONFIRMED, CANCELLED)
	   - Implementar paginação para voos com muitas reservas
	   - Ordenação por data de criação, nome do passageiro, etc.

	3. DADOS RETORNADOS:
	   - Lista de reservas com informações básicas
	   - Status de cada reserva
	   - Informações do passageiro (nome, documento)
	   - Número do assento
	   - Status do pagamento

	4. ESTATÍSTICAS:
	   - Total de reservas por status
	   - Ocupação atual do voo
	   - Receita total das reservas confirmadas
	   - Assentos disponíveis restantes

	5. CASOS ESPECIAIS:
	   - Não mostrar reservas canceladas por padrão
	   - Destacar reservas com pagamento pendente
	   - Mostrar reservas próximas do vencimento
	   - Permitir exportação para relatórios
*/
