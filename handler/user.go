package handler

import (
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/helper"
	"wishlist-wrangler-api/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func CreateUserEndpoint(dbClient *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		createUserDto := new(dto.CreateUserDto)

		if err := c.BodyParser(createUserDto); err != nil {
			return sendBadRequestResponse(c, err, "Failed to parse create user data")
		}

		isEmailValid := helper.DidUserEmailPassValidation(createUserDto.Email)

		if !isEmailValid {
			log.Warn("User submitted an invalid email address", createUserDto.Email)
			return sendBadRequestResponse(c, nil, "Email is not valid")
		}

		isDisplayNameValid := helper.DidUserDisplayNamePassValidation(createUserDto.DisplayName)

		if !isDisplayNameValid {
			log.Warn("User submitted an invalid email address", createUserDto.DisplayName)
			return sendBadRequestResponse(c, nil, "DisplayName is not valid")
		}

		_, err := repository.CreateUser(dbClient, *createUserDto)

		if err != nil {
			return sendInternalServerErrorResponse(c, err)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message":        "User Created",
			"loginRequestId": "TODO",
		})
	}
}
