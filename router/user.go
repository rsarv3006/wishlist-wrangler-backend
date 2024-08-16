package router

import (
	"wishlist-wrangler-api/ent"

	"github.com/gofiber/fiber/v2"
)

func setUpUserRoutesV1(api fiber.Router, _ *ent.Client) {
	_ = api.Group("/user")

}
