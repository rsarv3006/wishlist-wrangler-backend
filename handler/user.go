package handler

import (
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/repository"

	"github.com/gofiber/fiber/v2"
)

func DeleteUser(dbClient *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentUser := c.Locals("currentUser").(*ent.User)

		if err := repository.DeleteUser(dbClient, currentUser.ID); err != nil {
			return sendBadRequestResponse(c, err, "Failed to delete user")
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User deleted",
		})
	}
}
