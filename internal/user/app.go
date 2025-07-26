package user

import (
	"fmt"
	"os"

	"github.com/danielalmeidafarias/go-saga/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type UserApp struct {
	app *fiber.App
}

func NewUserApp() *UserApp {
	_ = godotenv.Load()

	db, err := pkg.InitializeDB(&User{})
	if err != nil {
		panic(err)
	}

	validator := pkg.NewValidator()

	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository, validator)
	userHandler := NewUserHandler(userService)

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

	userHandler.RegisterRoutes(app)

	return &UserApp{
		app: app,
	}
}

func (a *UserApp) Run() {
	if a.app != nil {
		a.app.Listen(fmt.Sprintf(":%s", os.Getenv("USER_APP_PORT")))
	}

	panic("app was not set up correctly")
}
