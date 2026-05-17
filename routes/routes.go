package routes

import (
	"github.com/bmfs-devs/sb_dashboard_be/handlers"
	"github.com/gofiber/fiber/v2"
)

func InitFiber() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello Fiber",
		})
	})
	remote := app.Group("/remote")
	remote.Post("/on-off", handlers.OnOffHandler)
	app.Get("/hello", handlers.HelloWorldHandler)
	app.Listen(":8080")
}
