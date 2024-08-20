package router

import (
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/handler"

	"github.com/gofiber/fiber/v2"
)

func setUpUserRoutesV1(api fiber.Router, dbClient *ent.Client) {
	user := api.Group("/user")

	user.Delete("/", handler.DeleteUser(dbClient))
}
