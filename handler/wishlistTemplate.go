package handler

import (
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/repository"

	"github.com/gofiber/fiber/v2"
)

func CreateWishlistTemplate(dbClient *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentUser := c.Locals("currentUser").(*ent.User)

		// TODO: Fully validate the request body
		createWishTemplateDto := new(dto.CreateWishlistTemplateDto)

		if err := c.BodyParser(createWishTemplateDto); err != nil {
			return sendBadRequestResponse(c, err, "Failed to parse create wishlist template data")
		}

		if len(createWishTemplateDto.Title) == 0 {
			return sendBadRequestResponse(c, nil, "Title is required")
		}

		wishlistTemplate, err := repository.CreateWishlistTemplate(dbClient, createWishTemplateDto, currentUser)

		if err != nil {
			return sendBadRequestResponse(c, err, "Failed to create wishlist template")
		}

		wishlistTemplateSections, err := repository.CreateWishlistTemplateSections(dbClient, createWishTemplateDto.TemplateSections, wishlistTemplate.ID)

		if err != nil {
			return sendBadRequestResponse(c, err, "Failed to create wishlist template sections")
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"wishlistTemplate":         wishlistTemplate,
			"wishlistTemplateSections": wishlistTemplateSections,
		})
	}
}
