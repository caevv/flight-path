package main

import (
	"github.com/caevv/volume/configs"
	"github.com/caevv/volume/internal/server"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	if err := configs.Load(); err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	s := server.New(fiber.New())

	if err := s.Start(); err != nil {
		log.Fatalf("failed to load config: %s", err)
	}
}
