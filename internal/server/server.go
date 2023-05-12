package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app    *fiber.App
	config *Config
}

func New(config *Config) *Server {
	app := fiber.New()
	initRoutes(app)

	return &Server{
		app:    app,
		config: config,
	}
}

func (s *Server) Run() {
	log.Fatal(s.app.Listen(fmt.Sprintf(":%d", s.config.Port)))
}
