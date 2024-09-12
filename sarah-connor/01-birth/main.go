package main

import (
	"fmt"
	"log"
	"os"

	"github.com/parakeet-nest/gollama"
)

func main() {
	ollamaUrl := os.Getenv("OLLAMA_URL")
	model := os.Getenv("LLM")

	systemContent := `You are Sarah Connor.`

	//userContent := `Give me a list of all the terminators models`
	//userContent := `Who is Sarah Connor?`
	//userContent := `Who is John Connor for you?`
	userContent := `What is your name? How can you save the world?`


	options := o.Options{
		Temperature: 0.5, // default (0.8)
		RepeatLastN: 2,   // default (64) the default value will "freeze" deepseek-coder
	}

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
