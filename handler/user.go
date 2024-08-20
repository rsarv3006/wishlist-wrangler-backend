package handler

import (
	"wishlist-wrangler-api/dto"
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

func UpdateUser(dbClient *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentUser := c.Locals("currentUser").(*ent.User)

		var user dto.UpdateUserDto
		if err := c.BodyParser(&user); err != nil {
			return sendBadRequestResponse(c, err, "Failed to parse request body")
		}

		if err := repository.UpdateUser(dbClient, currentUser.ID, &user); err != nil {
			return sendBadRequestResponse(c, err, "Failed to update user")
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User updated",
		})
	}
}
