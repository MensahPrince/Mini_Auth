package main

import (
	"log"

	"github.com/MensahPrince/mini_auth/db"
	"github.com/MensahPrince/mini_auth/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env file:", err)
	}

	db.Connect()

	app := fiber.New()

	routes.SetupAuthRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
