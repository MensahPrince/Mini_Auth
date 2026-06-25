package handlers

import (
	"github.com/MensahPrince/mini_auth/utils"
	"github.com/gofiber/fiber/v3"
)

// Root of API
func Base(c fiber.Ctx) error {

	status := utils.CheckDB()

	return c.Status(200).JSON(fiber.Map{
		"status": status,
		"server": "mini_auth",
	})
}
