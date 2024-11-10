package main

import (
	"03-memory/ui"
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

	brain := ""
	if len(os.Args) < 2 {
		// default brain
		brain = "v4"
	} else {
		brain = os.Args[1]
	}


	ctx := context.Background()
	errEnv := godotenv.Load(fmt.Sprintf("./data/brain-%s/.env", brain))

	ollamaUrl := os.Getenv("OLLAMA_HOST")
	model := os.Getenv("LLM")

	ui.Println("#ffc0c5", "🌍", ollamaUrl, "📕", model)

	client, errCli := api.ClientFromEnvironment()

	systemInstructionsFile, errInst := os.ReadFile(fmt.Sprintf("./data/brain-%s/instructions.md", brain))
	systemInstructions := string(systemInstructionsFile)

	personalityFile, errPers := os.ReadFile(fmt.Sprintf("./data/brain-%s/personality.md", brain))
	personality := string(personalityFile)

	// Configuration
	configFile, errConf := os.ReadFile(fmt.Sprintf("./data/brain-%s/settings.json", brain))
	var config map[string]interface{}
	errJsonConf := json.Unmarshal(configFile, &config)

	errorsList := errors.Join(errEnv, errCli, errInst, errPers, errConf, errJsonConf)
	if errorsList != nil {
		log.Fatal("😡:", errorsList)
	}

	ui.Println("#FFFF00", "📝 config:", config)

	// 🧠 Conversational memory
	memory := []api.Message{}

	for {
		question, _ := ui.Input("#008000", fmt.Sprintf("🤖 [%s] 🧠 (%s) ask me something> ", model, brain))

		if question == "bye" {
			break
		}

		// Prompt construction
		messages := []api.Message{
			{Role: "system", Content: systemInstructions},
			{Role: "system", Content: personality},
		}
		// 👋 🧠 Add memory
		messages = append(messages, memory...)
		// Add the new question
		messages = append(messages, api.Message{Role: "user", Content: question})

		req := &api.ChatRequest{
			Model:    model,
			Messages: messages,
			Options:  config,
		}

		answer := ""

		respFunc := func(resp api.ChatResponse) error {
			fmt.Print(resp.Message.Content)
			answer += resp.Message.Content
			return nil
		}

		err := client.Chat(ctx, req, respFunc)

		// 👋 🧠 Save the conversation in memory
		// Add the answer from the 🤖 agent to the memory
		
		memory = []api.Message{
			{Role: "user", Content: question},
			{Role: "assistant", Content: answer},
		}
		/*
		memory = append(
			memory,
			api.Message{Role: "user", Content: question},
			api.Message{Role: "assistant", Content: answer},
		)
		*/

		if err != nil {
			log.Fatal("😡:", err)
		}
		fmt.Println()
		fmt.Println()

	}

	/*
		- who are you?
		- who is Sarah Connor?

		🖐️
		- who is John Connor?
		- what is his job?

		- what is skynet?
		- give the list of all the terminators models

		🖐️
		- what is the T-800?
		- can you describe the T-1000 Series
	*/
}
