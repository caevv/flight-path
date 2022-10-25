package server

import (
	"encoding/json"
	"fmt"
	"github.com/caevv/volume/configs"
	"github.com/caevv/volume/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	fiber *fiber.App
}

func New(app *fiber.App) *Server {
	app.Use(recover.New())

	app.Post("/calculate", calculate)

	return &Server{
		fiber: app,
	}
}

func (s *Server) Start() error {
	return s.fiber.Listen(fmt.Sprintf(":%d", configs.Settings.ServerPort))
}

func calculate(c *fiber.Ctx) error {
	var paths []routes.FlightRoute

	if err := json.Unmarshal(c.Body(), &paths); err != nil {
		return fmt.Errorf("failed to parse body: %w", err)
	}

	flightRoute, err := routes.Calculate(paths)
	if err != nil {
		return fmt.Errorf("failed to calculate: %w", err)
	}

	return c.JSON(flightRoute)
}
