package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func sendBadRequestResponse(c *fiber.Ctx, err error, message string) error {
	log.Println(message)
	if err != nil {
		log.Println(err)
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": message,
		"error":   err.Error(),
	})
}

func sendInternalServerErrorResponse(c *fiber.Ctx, err error) error {
	if err != nil {
		log.Println(err)
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Internal Server Error",
		"error":   err,
	})
}

func sendNotFoundResponse(c *fiber.Ctx, err error) error {
	if err != nil {
		log.Println(err)
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Not Found",
		"error":   err,
	})
}
