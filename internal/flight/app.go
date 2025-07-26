package flight

import (
	"fmt"
	"os"

	"github.com/danielalmeidafarias/go-saga/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type FlightApp struct {
	app *fiber.App
}

func NewFlightApp() *FlightApp {
	_ = godotenv.Load()

	db, err := pkg.InitializeDB(&Flight{}, &Seat{})
	if err != nil {
		panic(err)
	}

	validator := pkg.NewValidator()

	flightRepository := NewFlightRepository(db)
	flightService := NewFlightService(flightRepository, validator)
	flightHandler := NewFlightHandler(flightService)

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

	flightHandler.RegisterRoutes(app)

	return &FlightApp{
		app: app,
	}
}

func (a *FlightApp) Run() {
	if a.app != nil {
		a.app.Listen(fmt.Sprintf(":%s", os.Getenv("FLIGHT_APP_PORT")))
	}

	panic("app was not set up correctly")
}
