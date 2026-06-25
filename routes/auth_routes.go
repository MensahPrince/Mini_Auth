package routes

import (
	"github.com/MensahPrince/mini_auth/handlers"
	"github.com/gofiber/fiber/v3"
)

func SetupAuthRoutes(app *fiber.App) {
	app.Get("/", handlers.Base)
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
}
