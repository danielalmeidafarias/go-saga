package flight

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type FlightHandler struct {
	flightService *FlightService
}

func NewFlightHandler(
	flightService *FlightService,
) *FlightHandler {
	return &FlightHandler{
		flightService: flightService,
	}
}

func (h *FlightHandler) RegisterRoutes(app *fiber.App) {
	flightGroup := app.Group("/flights")

	flightGroup.Post("/", h.CreateFlight)
	flightGroup.Get("/:id", h.GetFlight)
	flightGroup.Put("/:id", h.UpdateFlight)
	flightGroup.Delete("/:id", h.DeleteFlight)

	flightGroup.Post("/:id/seats", h.CreateFlightSeat)
	flightGroup.Get("/seats/:id", h.GetFlightSeat)
	flightGroup.Put("/seats/:id", h.UpdateFlightSeat)
}

func (h *FlightHandler) CreateFlight(c *fiber.Ctx) error {
	var params CreateFlightParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	response, err := h.flightService.CreateFlight(c.Context(), params)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(response)
}

func (h *FlightHandler) GetFlight(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Flight ID is required",
		})
	}

	response, err := h.flightService.GetFlight(c.Context(), id)
	if err != nil {
		if err.Error() == "flight not found" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Flight not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(response)
}

func (h *FlightHandler) UpdateFlight(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Flight ID is required",
		})
	}

	var params UpdateFlightParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	flight, err := h.flightService.UpdateFlight(c.Context(), id, params)
	if err != nil {
		if err.Error() == "flight not found" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Flight not found",
			})
		}
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(flight)
}

func (h *FlightHandler) DeleteFlight(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Flight ID is required",
		})
	}

	err := h.flightService.DeleteFlight(c.Context(), id)
	if err != nil {
		if err.Error() == "flight not found" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Flight not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Flight deleted successfully",
	})
}

func (h *FlightHandler) CreateFlightSeat(c *fiber.Ctx) error {
	flightID := c.Params("id")

	if flightID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Flight ID is required",
		})
	}

	priceStr := c.FormValue("price")
	if priceStr == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Price is required",
		})
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid price format",
		})
	}

	seat, err := h.flightService.CreateFlightSeat(c.Context(), flightID, price)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(seat)
}

func (h *FlightHandler) GetFlightSeat(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Seat ID is required",
		})
	}

	seat, err := h.flightService.GetFlightSeat(c.Context(), id)
	if err != nil {
		if err.Error() == "seat not found" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Seat not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(seat)
}

func (h *FlightHandler) UpdateFlightSeat(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Seat ID is required",
		})
	}

	var params UpdateSeatInput

	if err := c.BodyParser(&params); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	seat, err := h.flightService.UpdateFlightSeat(c.Context(), id, params)
	if err != nil {
		if err.Error() == "seat not found" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Seat not found",
			})
		}
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(seat)
}
