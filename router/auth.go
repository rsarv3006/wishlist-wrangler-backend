package router

import (
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/handler"

	"github.com/gofiber/fiber/v2"
)

func setUpAuthRoutesV1(api fiber.Router, dbClient *ent.Client) {
	auth := api.Group("/auth")

	auth.Post("/signup", handler.SignUpEndpoint(dbClient))
	auth.Post("/login", handler.LoginEndpoint(dbClient))
}
