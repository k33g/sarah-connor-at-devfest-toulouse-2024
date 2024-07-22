package sidekicks

import (
	"fmt"
	"os"
	"github.com/parakeet-nest/gollama"
)

func Translator(content string) {

	fmt.Println()
	fmt.Println()
	fmt.Println("🇫🇷 translation:")

	ollamaUrl := os.Getenv("OLLAMA_URL")

	queryForTranslation := gollama.Query{
		Model: "qwen2:1.5b",
		Messages: []gollama.Message{
			{Role: "system", Content: "You are a useful AI agent, your are an expert with translation from English to French."},
			{Role: "user", Content: "Translate this content in French:" + content},
		},
		Options: gollama.Options{
			Temperature: 0.5,
			RepeatLastN: 2,
		},
	}

	_, _ = gollama.ChatStream(ollamaUrl, queryForTranslation,
		func(answer gollama.Answer) error {
			fmt.Print(answer.Message.Content)
			return nil
		},
	)

	fmt.Println()

}
