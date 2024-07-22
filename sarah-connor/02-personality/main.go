package main

import (
	"02-personality/sidekicks"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/parakeet-nest/gollama"
)

func main() {
	ollamaUrl := os.Getenv("OLLAMA_URL")
	model := os.Getenv("LLM")

	systemInstructions := os.Getenv("SYSTEM_INSTRUCTIONS")
	personality := os.Getenv("PERSONALITY")

	userMessage := os.Getenv("USER_MESSAGE")

	temperature, _ := strconv.ParseFloat(os.Getenv("TEMPERATURE"), 64)
	repeatLastN, _ := strconv.Atoi(os.Getenv("REPEAT_LAST_N"))
	topK, _ := strconv.Atoi(os.Getenv("TOP_K"))
	topP, _ := strconv.ParseFloat(os.Getenv("TOP_P"), 64)

	options := gollama.Options{
		Temperature: temperature, // default (0.8)
		RepeatLastN: repeatLastN, // default (64) the default value will "freeze" deepseek-coder
		TopK:        topK,
		TopP:        topP,
	}

	query := gollama.Query{
		Model: model,
		Messages: []gollama.Message{
			{Role: "system", Content: systemInstructions},
			{Role: "system", Content: personality},
			{Role: "user", Content: userMessage},
		},
		Options: options,
	}

	response, err := gollama.ChatStream(ollamaUrl, query,
		func(answer gollama.Answer) error {
			fmt.Print(answer.Message.Content)
			return nil
		},
	)

	if err != nil {
		log.Fatal("😡:", err)
	}

	sidekicks.Translator(response.Message.Content)

}
