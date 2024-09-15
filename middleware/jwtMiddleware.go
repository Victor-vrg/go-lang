package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/dgrijalva/jwt-go"
	"os"
)

func JWTMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	claims := &struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	return c.Next()
}
