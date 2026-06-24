package main

import (
	"log"
	
	
	"github.com/MensahPrince/mini_auth/db"
	"github.com/MensahPrince/mini_auth/types"
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
	app.Post("/add", func(c fiber.Ctx) error {

		var req types.USER

		//Check for JSON body
		if err := c.Bind().Body(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid Request",
			})
		}

		//Write to Database
		_, err := database.Exec(
			"INSERT INTO users (name, email, password) VALUES (?,?,?,?)",
			req.Name,
			req.Email,
			req.Password,
		)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message":  "Success",
			"name":     req.Name,
			"email":    req.Email,
			"password": req.Password,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
