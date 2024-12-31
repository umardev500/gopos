package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"gitub.com/umardev500/gopos/pkg/auth"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the token from the header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Split the header to get the token (it should be in the format "Bearer <token>")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Verify the token
		claims, err := auth.VerifyJWT(tokenString)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Add the claims to the context
		c.Locals("claims", claims)

		// Pass the request to the next middleware
		return c.Next()
	}
}
