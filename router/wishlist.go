package router

import (
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/handler"
	"wishlist-wrangler-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func setUpWishlistRoutesV1(api fiber.Router, dbClient *ent.Client) {
	wishlist := api.Group("/wishlist")
	wishlist.Use(middleware.IsExpired())

	wishlist.Post("/", handler.CreateWishlist(dbClient))
	wishlist.Get("/ByUser", handler.GetWishlistsByUser(dbClient))
	wishlist.Delete("/:id", handler.DeleteWishlist(dbClient))
}
