package sidekicks

import (
	"fmt"
	"os"
	"github.com/parakeet-nest/gollama"
	"github.com/parakeet-nest/gollama/enums/option"

)

func Translator(content string) {

	fmt.Println()
	fmt.Println()
	fmt.Println("🇫🇷 translation:")

	ollamaUrl := os.Getenv("OLLAMA_URL")

	options := gollama.SetOptions(map[string]interface{}{
		option.Temperature: 0.5,
		option.RepeatLastN: 2,
	})

	queryForTranslation := gollama.Query{
		Model: "qwen2:1.5b",
		Messages: []gollama.Message{
			{Role: "system", Content: "You are a useful AI agent, your are an expert with translation from English to French."},
			{Role: "user", Content: "Translate this content in French:" + content},
		},
		Options: options,
	}

	_, _ = gollama.ChatStream(ollamaUrl, queryForTranslation,
		func(answer gollama.Answer) error {
			fmt.Print(answer.Message.Content)
			return nil
		},
	)

	fmt.Println()

}
