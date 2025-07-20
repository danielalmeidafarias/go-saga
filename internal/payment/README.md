# Payment Service

Serviço de pagamentos que simula um provedor de pagamento externo com comportamento assíncrono.

## Funcionalidades

- **CreatePayment**: Cria um novo pagamento com status "pending"
- **GetPayment**: Obtém informações de um pagamento (rota de notificação)
- **ProcessPayment**: Processa um pagamento de forma assíncrona com simulação de falhas

## Características do Serviço

### Comportamento Assíncrono
- O processamento de pagamentos é feito em background
- Simula tempo de processamento entre 1-5 segundos
- Retorna imediatamente com status HTTP 202 (Accepted)

### Simulação de Falhas
- 30% de chance de falha no processamento
- Status possíveis: `pending`, `processed`, `failed`

### Estados do Pagamento
- **pending**: Pagamento criado mas não processado
- **processed**: Pagamento processado com sucesso
- **failed**: Falha no processamento do pagamento

## API Endpoints

### POST /payments
Cria um novo pagamento
```json
{
  "description": "Flight booking payment",
  "value": 299.99
}
```

### GET /payments/:id
Obtém informações do pagamento (rota de notificação)
```json
{
  "paymentUUID": "123e4567-e89b-12d3-a456-426614174000",
  "description": "Flight booking payment",
  "value": 299.99,
  "status": "processed"
}
```

### POST /payments/:id/process
Inicia o processamento do pagamento (assíncrono)
```json
{
  "message": "Payment processing started",
  "paymentUUID": "123e4567-e89b-12d3-a456-426614174000"
}
```

## Configuração

O serviço utiliza as seguintes variáveis de ambiente:
- `PAYMENT_APP_PORT`: Porta do serviço (padrão: 3002)
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`: Configurações do banco de dados

## Executando o Serviço

```bash
go run cmd/payment_service/main.go
```

O serviço estará disponível em `http://localhost:3002`
