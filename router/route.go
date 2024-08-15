package router

import (
	"os"
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

}

func setUpWellKnownRoutes(app *fiber.App) {
	wellKnown := app.Group("/.well-known")
	wellKnown.Get("/apple-app-site-association", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json; charset=utf-8")

		content, err := os.ReadFile("apple-app-site-association.json")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error reading AASA file")
		}

		return c.Send(content)
	})
}

func setUpShareRedirectRoutes(app *fiber.App) {
	app.Get("/share", func(c *fiber.Ctx) error {
		return c.Redirect("https://apps.apple.com/us/app/basketbuddy/id6446040498")
	})
}

func setUpUserRoutesV1(api fiber.Router, _ *ent.Client) {
	_ = api.Group("/user")

}

func setUpAuthRoutesV1(api fiber.Router, dbClient *ent.Client) {
	auth := api.Group("/auth")

	auth.Post("/code/:code", handler.CreateUserEndpoint(dbClient))
}
