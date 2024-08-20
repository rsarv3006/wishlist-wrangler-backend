package router

import (
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/handler"
	"wishlist-wrangler-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func setUpUserRoutesV1(api fiber.Router, dbClient *ent.Client) {
	user := api.Group("/user")
	user.Use(middleware.IsExpired())

	user.Delete("/", handler.DeleteUser(dbClient))
}
