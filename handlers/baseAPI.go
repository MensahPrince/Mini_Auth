package handlers

import (
	"log"

	"github.com/MensahPrince/mini_auth/db"
	"github.com/MensahPrince/mini_auth/utils"
	"github.com/gofiber/fiber/v3"
)

// Root of API
func Base(c fiber.Ctx) error {
	_, err := db.Connect()
	status := utils.CheckDB()
	if err != nil {
		log.Fatal("Failed to Initialize", err)
	}
	return c.Status(200).JSON(fiber.Map{
		"status": status,
		"server": "mini_auth",
	})
}
