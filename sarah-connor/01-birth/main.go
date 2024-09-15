package main

import (
	"fmt"
	"log"
	"os"

	"github.com/parakeet-nest/gollama"
	"github.com/parakeet-nest/gollama/enums/option"

)

func main() {
	ollamaUrl := os.Getenv("OLLAMA_URL")
	model := os.Getenv("LLM")

	systemContent := `You are Sarah Connor.`

	//userContent := `Give me a list of all the terminators models`
	//userContent := `Who is Sarah Connor?`
	//userContent := `Who is John Connor for you?`
	userContent := `What is your name? How can you save the world?`


	options := gollama.SetOptions(map[string]interface{}{
		option.Temperature: 0.5,
		option.RepeatLastN: 2,
	})

	query := gollama.Query{
		Model: model,
		Messages: []gollama.Message{
			{Role: "system", Content: systemContent},
			{Role: "user", Content: userContent},
		},
		Options: options,
	}

	_, err := gollama.ChatStream(ollamaUrl, query,
		func(answer gollama.Answer) error {
			fmt.Print(answer.Message.Content)
			return nil
		})

	if err != nil {
		log.Fatal("😡:", err)
	}
}
