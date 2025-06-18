package routes

import (
	"prompt-server/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupCoreRouter(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	v1.Post("/chat", handlers.ChatHandler())
}
