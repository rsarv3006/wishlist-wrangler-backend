package router

import (
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/handler"
	"wishlist-wrangler-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func setUpWishlistTemplateRoutesV1(api fiber.Router, dbClient *ent.Client) {
	wishlistTemplate := api.Group("/wishlistTemplate")
	wishlistTemplate.Use(middleware.IsExpired())

	wishlistTemplate.Post("/", handler.CreateWishlistTemplate(dbClient))
	wishlistTemplate.Get("/ByUser", handler.GetWishlistTemplatesForUser(dbClient))
	wishlistTemplate.Delete("/:templateId", handler.DeleteWishlistTemplate(dbClient))
}
