package handlers

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/genai"
)

// incoming request payload
type chatRequest struct {
	Prompt string `json:"prompt"`
}

// outgoing response payload
type chatResponse struct {
	Response string `json:"response"`
}

func ChatHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// 1. Parse JSON body.
		var req chatRequest
		if err := ctx.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
		}
		if req.Prompt == "" {
			return fiber.NewError(fiber.StatusBadRequest, "`prompt` cannot be empty")
		}


		//MOVE THIS LATER
		apiKey := os.Getenv("GEMINI_API_KEY")
		client, clientErr := genai.NewClient(ctx.Context(), &genai.ClientConfig{
			APIKey:  apiKey,
			Backend: genai.BackendGeminiAPI,
		})
		if clientErr != nil {
			log.Fatal(clientErr)
		}

		result, promptErr := client.Models.GenerateContent(
			ctx.Context(),
			"gemini-2.5-flash",
			genai.Text("Explain how AI works in a few words"),
			nil,
		)

		if promptErr != nil {
			// OpenAI returned an error — tell the caller.
			return fiber.NewError(fiber.StatusBadGateway, clientErr.Error())
		}

		// 3. Send the assistant’s first choice back as JSON.
		return ctx.JSON(result)
	}
}
