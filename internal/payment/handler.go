package payment

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type PaymentHandler struct {
	paymentService *PaymentService
}

func NewPaymentHandler(
	paymentService *PaymentService,
) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
	}
}

func (h *PaymentHandler) RegisterRoutes(app *fiber.App) {
	paymentGroup := app.Group("/payments")

	paymentGroup.Post("/", h.CreatePayment)
	paymentGroup.Get("/:id", h.GetPayment)              // Rota de notificação
	paymentGroup.Post("/:id/process", h.ProcessPayment) // Rota para processar pagamento
}

func (h *PaymentHandler) CreatePayment(c *fiber.Ctx) error {
	var params CreatePaymentParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	response, err := h.paymentService.CreatePayment(c.Context(), params)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(response)
}

func (h *PaymentHandler) GetPayment(c *fiber.Ctx) error {
	id := c.Params("id")

	response, err := h.paymentService.GetPayment(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(response)
}

func (h *PaymentHandler) ProcessPayment(c *fiber.Ctx) error {
	id := c.Params("id")

	// Inicia o processamento em uma goroutine para simular comportamento assíncrono
	go func() {
		h.paymentService.ProcessPayment(c.Context(), id)
	}()

	return c.Status(http.StatusAccepted).JSON(fiber.Map{
		"message":     "Payment processing started",
		"paymentUUID": id,
	})
}
