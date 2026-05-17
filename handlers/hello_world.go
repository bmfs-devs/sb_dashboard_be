package handlers

import "github.com/gofiber/fiber/v2"

// Handler only for parse input and output, no business logic may appear here
func HelloWorldHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello World",
	})
}
