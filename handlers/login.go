package handlers

import (
	"github.com/MensahPrince/mini_auth/db"
	"github.com/MensahPrince/mini_auth/types"
	"github.com/MensahPrince/mini_auth/utils"
	"github.com/gofiber/fiber/v3"
)

func Login(c fiber.Ctx) error {
	database, err := db.Connect()

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Failed to Initialize",
		})
	}

	var req types.USERLOGIN

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid JSON request",
		})
	}

	var hashedPassword string

	err = database.QueryRow(
		"SELECT password FROM users WHERE email = ?",
		req.Email,
	).Scan(&hashedPassword)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	//corrected password check
	if !utils.BcryptCompareHash(req.Password, hashedPassword) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
	})

}
