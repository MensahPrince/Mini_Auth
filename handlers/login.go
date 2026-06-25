package handlers

import (
	"github.com/MensahPrince/mini_auth/db"
	"github.com/MensahPrince/mini_auth/types"
	"github.com/gofiber/fiber/v3"
)

func Login(c fiber.Ctx) error {
	database, err := db.Connect()

	var state bool

	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"error": "Failed to Initialize",
		})
	}

	var req types.USERLOGIN

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid JSON request",
		})
	}

	err = database.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)",
		req.Email,
	).Scan(&state)
	if &state == false {
		c.JSON(fiber.Map{
			"response" : "User does not exist",
		})
	}

   var hashedPassword string

   err := database.QueryRow(
    "SELECT password FROM users WHERE email = ?",
    req.Email,
   ).Scan(&hashedPassword)


		

	var ComparedHashResult = utils.BcryptCompareHash(types.USERLOGIN.Password, hashedPassword)
