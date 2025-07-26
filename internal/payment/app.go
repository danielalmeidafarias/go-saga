package payment

import (
	"fmt"
	"os"

	"github.com/danielalmeidafarias/go-saga/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type PaymentApp struct {
	app *fiber.App
}

func NewPaymentApp() *PaymentApp {
	_ = godotenv.Load()

	db, err := pkg.InitializeDB(&Payment{})
	if err != nil {
		panic(err)
	}
	validator := pkg.NewValidator()

	paymentRepository := NewPaymentRepository(db)
	paymentService := NewPaymentService(paymentRepository, validator)
	paymentHandler := NewPaymentHandler(paymentService)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	paymentHandler.RegisterRoutes(app)

	return &PaymentApp{
		app: app,
	}
}

func (a *PaymentApp) Run() {
	if a.app != nil {
		a.app.Listen(fmt.Sprintf(":%s", os.Getenv("PAYMENT_APP_PORT")))
	}

	panic("app was not set up correctly")
}
