package authutils

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func GenerateSymmetricJWT(c *fiber.Ctx) (string, error) {
	if len(jwtKey) == 0 {
		return "", fiber.NewError(fiber.StatusInternalServerError, "Failed to load ENV variables.")
	}

	t := jwt.New(jwt.SigningMethodHS256)
	return t.SignedString(jwtKey)
}
