package flight

import (
	"fmt"
	"os"

	"github.com/danielalmeidafarias/go-saga/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type FlightApp struct {
	app *fiber.App
}

func NewFlightApp() *FlightApp {
	_ = godotenv.Load()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	if err := db.AutoMigrate(&Flight{}, &Seat{}); err != nil {
		panic("failed to migrate database: " + err.Error())
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
