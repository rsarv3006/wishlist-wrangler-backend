package middleware

import (
	"strings"
	"wishlist-wrangler-api/auth"

	"github.com/gofiber/fiber/v2"
)

func IsExpired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !strings.Contains(c.Get("Authorization"), "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
				"error":   "No token provided",
			})
		}

		token := strings.Split(c.Get("Authorization"), "Bearer ")[1]
		currentUser, err := auth.ValidateToken(token, c)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
				"error":   err,
			})
		}

		c.Locals("currentUser", currentUser)

		defer func() {
			currentUser = nil
		}()

		return c.Next()
	}
}
