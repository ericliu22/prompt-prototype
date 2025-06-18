package prompt

import (
	"log"
	"os"
)

var basePrompt string

func LoadGPTRules(path string) {
    b, err := os.ReadFile(path)
    if err != nil {
        log.Fatalf("cannot read %s: %v", path, err)
    }
    basePrompt = string(b)
}

func InjectPrompt(userPrompt string) string {
	    return basePrompt + "\n\nUser: " + userPrompt + "\nAI:"
}
