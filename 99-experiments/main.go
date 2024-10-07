package main

import (
	"02-personality/ui"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ollama/ollama/api"
)

func main() {
	brain := "v1"
	//brain := "v2"
	//brain := "v3"
	//brain := "v4"

	ctx := context.Background()
	errEnv := godotenv.Load(fmt.Sprintf("./data/brain-%s/.env", brain))

	ollamaUrl := os.Getenv("OLLAMA_HOST")
	model := os.Getenv("LLM")

	ui.Println("🌍", ollamaUrl, "📕", model)

	client, errCli := api.ClientFromEnvironment()

	systemInstructionsFile, errInst := os.ReadFile(fmt.Sprintf("./data/brain-%s/instructions.md", brain))
	systemInstructions := string(systemInstructionsFile)

	personalityFile, errPers := os.ReadFile(fmt.Sprintf("./data/brain-%s/personality.md", brain))
	personality := string(personalityFile)

	// Configuration
	configFile, errConf := os.ReadFile(fmt.Sprintf("./data/brain-%s/options.json", brain))
	var config map[string]interface{}
	errJsonConf := json.Unmarshal(configFile, &config)

	errorsList := errors.Join(errEnv, errCli, errInst, errPers, errConf, errJsonConf)
	if errorsList != nil {
		log.Fatal("😡:", errorsList)
	}

	fmt.Println("📝 config:", config)

	for {
		question, _ := ui.Input("#008000", fmt.Sprintf("🤖 [%s] 🧠 (%s) ask me something> ", model, brain))

		if question == "bye" {
			break
		}

		// Prompt construction
		messages := []api.Message{
			{Role: "system", Content: systemInstructions},
			{Role: "system", Content: personality},
			{Role: "user", Content: question},
		}

		req := &api.ChatRequest{
			Model:    model,
			Messages: messages,
			Options:  config,
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
		- who is John Connor?
		- what is skynet?
		- give the list of all the terminators models
	*/
}
