package server

import (
	"authService/controllers"
	"authService/server/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetPublicRoutes(app *fiber.App) {
	app.Post("/users", controllers.SignUp)
	app.Post("/login", controllers.Login)
	app.Get("/validate", controllers.Validate)
}

func SetGlobalMiddleware(app *fiber.App) {
	app.Use(middleware.ValidateRateLimit)
}
