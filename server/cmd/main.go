package main

import (
	"log"
	"prompt-server/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	routes.SetupCoreRouter(app)

	log.Fatal(app.Listen(":3000"))
}
