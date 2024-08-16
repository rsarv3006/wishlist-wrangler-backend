package router

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

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
