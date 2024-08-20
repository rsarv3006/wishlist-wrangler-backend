package handler

import (
	"context"
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/helper"
	"wishlist-wrangler-api/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateWishlistTemplate(dbClient *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentUser := c.Locals("currentUser").(*ent.User)

		createWishTemplateDto := new(dto.CreateWishlistTemplateDto)

		if err := c.BodyParser(createWishTemplateDto); err != nil {
			return sendBadRequestResponse(c, err, "Failed to parse create wishlist template data")
		}

		if !helper.DidCreateWishlistTemplateDtoPassValidation(*createWishTemplateDto) {
			return sendBadRequestResponse(c, nil, "Failed to validate create wishlist template data")
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

func GetWishlistTemplatesForUser(dbClient *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentUser := c.Locals("currentUser").(*ent.User)

		wishlistTemplates, err := repository.GetWishlistTemplatesByUser(dbClient, currentUser.ID)

		if err != nil {
			return sendBadRequestResponse(c, err, "Failed to get wishlist templates")
		}

		templateIds := make([]uuid.UUID, 0, len(wishlistTemplates))
		for _, template := range wishlistTemplates {
			templateIds = append(templateIds, template.ID)
		}

		sections, err := repository.GetWishlistTemplateSectionsForTemplateArray(dbClient, templateIds)

		if err != nil {
			return sendBadRequestResponse(c, err, "Failed to get wishlist template sections")
		}

		sectionsByTemplateId := matchSectionsToTemplates(sections, templateIds)

		wishlistTemplatesAndSections := make([]*dto.WishlistTemplateAndSectionsReturnDto, 0, len(wishlistTemplates))
		for _, template := range wishlistTemplates {
			sections := sectionsByTemplateId[template.ID]
			wishlistTemplatesAndSections = append(wishlistTemplatesAndSections, &dto.WishlistTemplateAndSectionsReturnDto{
				WishlistTemplate:         template,
				WishlistTemplateSections: sections,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"wishlistTemplates": wishlistTemplatesAndSections,
		})
	}
}

func matchSectionsToTemplates(sections []*ent.WishlistTemplateSection, templateIds []uuid.UUID) map[uuid.UUID][]*ent.WishlistTemplateSection {
	result := make(map[uuid.UUID][]*ent.WishlistTemplateSection)

	for _, id := range templateIds {
		result[id] = []*ent.WishlistTemplateSection{}
	}

	for _, section := range sections {
		templateId := section.WishlistTemplateID
		if _, ok := result[templateId]; ok {
			result[templateId] = append(result[templateId], section)
		}
	}

	return result
}

func DeleteWishlistTemplate(dbClient *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentUser := c.Locals("currentUser").(*ent.User)

		templateId, err := uuid.Parse(c.Params("templateId"))

		if err != nil {
			return sendBadRequestResponse(c, err, "Failed to parse template id")
		}

		wishlistsUsingTemplate, err := repository.GetWishlistsByTemplate(dbClient, templateId)

		if err != nil {
			return err
		}

		if len(wishlistsUsingTemplate) > 0 {
			return sendBadRequestResponse(c, nil, "Template is in use by one or more wishlists")
		}

		if err := repository.DeleteWishlistTemplate(dbClient, templateId, currentUser.ID); err != nil {
			return sendBadRequestResponse(c, err, "Failed to delete wishlist template")
		}

		if err := repository.DeleteTemplateSectionsForTemplate(context.Background(), dbClient, templateId); err != nil {
			return sendBadRequestResponse(c, err, "Failed to delete wishlist template sections")
		}

		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"message": "Wishlist template deleted",
		})
	}
}
