package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Port string
}

func (s Server) Start() error {
	app := fiber.New()
	SetPublicRoutes(app)
	SetGlobalMiddleware(app)
	return app.Listen(fmt.Sprintf(":%s", s.Port))
}
