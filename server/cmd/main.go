package main

import (
	"log"
	"prompt-server/internal/core/prompt"
	"prompt-server/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	prompt.LoadGPTRules("internal/models/gpt_rules.txt")

	routes.SetupCoreRouter(app)

	log.Fatal(app.Listen(":3000"))
}
