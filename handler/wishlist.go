package handler

import (
	"wishlist-wrangler-api/dto"
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateWishlist(dbClient *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentUser := c.Locals("currentUser").(*ent.User)

		var createWishlistDto dto.CreateWishlistDto
		if err := c.BodyParser(&createWishlistDto); err != nil {
			return sendBadRequestResponse(c, err, "Failed to parse request body")
		}

		wishlist, err := repository.CreateWishlist(dbClient, createWishlistDto, currentUser)

		if err != nil {
			return sendBadRequestResponse(c, err, "Failed to create wishlist")
		}

		sections, err := repository.CreateWishlistSections(dbClient, createWishlistDto.Sections, wishlist.ID)

		if err != nil {
			return sendBadRequestResponse(c, err, "Failed to create wishlist sections")
		}

		// TODO: return a response with the created wishlist and sections
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message":  "Wishlist created",
			"wishlist": wishlist,
			"sections": sections,
		})
	}
}

func GetWishlistsByUser(dbClient *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentUser := c.Locals("currentUser").(*ent.User)
		wishlists, err := repository.GetWishlistsByUser(dbClient, currentUser.ID)

		if err != nil {
			return sendBadRequestResponse(c, err, "Failed to get wishlists")
		}

		wishlistIds := make([]uuid.UUID, 0, len(wishlists))
		for _, wishlist := range wishlists {
			wishlistIds = append(wishlistIds, wishlist.ID)
		}

		sections, err := repository.GetWishlistSectionsForWishlistArray(dbClient, wishlistIds)

		if err != nil {
			return sendBadRequestResponse(c, err, "Failed to get wishlist sections")
		}

		sectionsByWishlistId := matchSectionsToWishlists(sections, wishlistIds)

		wishlistsAndSections := make([]*dto.WishlistAndSectionsReturnDto, 0, len(wishlists))
		for _, wishlist := range wishlists {
			sections := sectionsByWishlistId[wishlist.ID]
			wishlistsAndSections = append(wishlistsAndSections, &dto.WishlistAndSectionsReturnDto{
				Wishlist:         wishlist,
				WishlistSections: sections,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"wishlists": wishlistsAndSections,
		})
	}
}

func DeleteWishlist(dbClient *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentUser := c.Locals("currentUser").(*ent.User)
		wishlistId := c.Params("id")

		wishlistUuid, err := uuid.Parse(wishlistId)
		if err != nil {
			return sendBadRequestResponse(c, err, "Invalid wishlist id")
		}

		wishlist, err := repository.GetWishlistById(dbClient, wishlistUuid)

		if err != nil {
			return sendNotFoundResponse(c, err)
		}

		if wishlist.CreatorID != currentUser.ID {
			return sendBadRequestResponse(c, err, "You are not the owner of this wishlist")
		}

		if err := repository.DeleteWishlist(dbClient, wishlistUuid, currentUser.ID); err != nil {
			return sendBadRequestResponse(c, err, "Failed to delete wishlist")
		}

		if err := repository.DeleteWishlistSectionsByWishlistId(dbClient, wishlist.ID); err != nil {
			return sendInternalServerErrorResponse(c, err)
		}

		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"message": "Wishlist deleted",
		})
	}
}

func matchSectionsToWishlists(sections []*ent.WishlistSection, wishlistIds []uuid.UUID) map[uuid.UUID][]*ent.WishlistSection {
	result := make(map[uuid.UUID][]*ent.WishlistSection)

	for _, id := range wishlistIds {
		result[id] = []*ent.WishlistSection{}
	}

	for _, section := range sections {
		wishlistId := section.WishlistID
		if _, ok := result[wishlistId]; ok {
			result[wishlistId] = append(result[wishlistId], section)
		}
	}

	return result
}
