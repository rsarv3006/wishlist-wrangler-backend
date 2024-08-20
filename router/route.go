package router

import (
	"wishlist-wrangler-api/ent"
	"wishlist-wrangler-api/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, dbClient *ent.Client) {
	setUpWellKnownRoutes(app)
	setUpShareRedirectRoutes(app)

	api := app.Group("/api", logger.New())
	api.Get("/health", handler.HealthEndpoint())
	apiV1 := api.Group("/v1")
	setUpUserRoutesV1(apiV1, dbClient)
	setUpAuthRoutesV1(apiV1, dbClient)
	setUpWishlistTemplateRoutesV1(apiV1, dbClient)
	setUpWishlistRoutesV1(apiV1, dbClient)
}
