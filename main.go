package main

import (
	"log"

	"github.com/MensahPrince/mini_auth/db"
	"github.com/MensahPrince/mini_auth/utils"
	"github.com/gofiber/fiber/v3"
)

func main() {
	database, err := db.Connect()

	status := utils.CheckDB()
	if err != nil {
		log.Fatal("Failed to Initialize", err)
	}

	app := fiber.New()

	//Root of API
	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status": status,
			"server": "mini_auth",
		})
	})

	//Add a new User.
}
