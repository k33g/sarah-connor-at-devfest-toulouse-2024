package main

import (
	"01-birth/ui" // tools
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ollama/ollama/api"
)

func main() {
	
	ctx := context.Background()
	errEnv := godotenv.Load()

	ollamaUrl := os.Getenv("OLLAMA_HOST")
	model := os.Getenv("LLM")

	ui.Println("🌍", ollamaUrl, "📕", model)

	client, errCli := api.ClientFromEnvironment()

	errorsList := errors.Join(errEnv, errCli)
	if errorsList != nil {
		log.Fatal("😡:", errorsList)
	}

	systemInstructions := `You a useful AI agent.`

	for {
		question, _ := ui.Input("#008000", fmt.Sprintf("🤖 [%s] ask me something> ", model))

		if question == "bye" {
			break
		}

		// Prompt construction
		messages := []api.Message{
			{Role: "system", Content: systemInstructions},
			{Role: "user", Content: question},
		}

		req := &api.ChatRequest{
			Model:    model,
			Messages: messages,
			Options:  map[string]interface{}{},
		}

		respFunc := func(resp api.ChatResponse) error {
			fmt.Print(resp.Message.Content)
			return nil
		}

		err := client.Chat(ctx, req, respFunc)
		if err != nil {
			log.Fatal("😡:", err)
		}
		fmt.Println()
		fmt.Println()

	}

	/*
		- who are you?
		- who is Sarah Connor?
		- who is Sarah Connor in The Terminator?
		- who is John Connor for you?
		- who is John Connor in Terminator?
		- give the list of all the terminators models
		- what is skynet?
	*/
}
