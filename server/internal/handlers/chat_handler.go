package handlers

import (
	"context"
	"log"
	"os"
	"prompt-server/internal/core/prompt"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/genai"
)

// incoming request payload
type chatRequest struct {
	Prompt string `json:"prompt"`
	Model  string `json:"model,omitempty"`
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

		client, clientErr := genai.NewClient(context.Background(), &genai.ClientConfig{
			APIKey:  apiKey,
			Backend: genai.BackendGeminiAPI,
		})
		if clientErr != nil {
			log.Fatal(clientErr)
		}

		engineeredPrompt := prompt.InjectPrompt(req.Prompt)

		result, promptErr := client.Models.GenerateContent(
			context.Background(),
			"models/gemini-2.5-flash",
			genai.Text(engineeredPrompt),
			nil,
		)

		if promptErr != nil {
			// OpenAI returned an error — tell the caller.
			return fiber.NewError(fiber.StatusBadGateway, promptErr.Error())
		}

		// 3. Send the assistant’s first choice back as JSON.
		return ctx.JSON(result)
	}
}
