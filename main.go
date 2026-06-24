package main

import (
	"log"

	"github.com/MensahPrince/mini_auth/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {

	app := fiber.New()

	routes.SetupAuthRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
