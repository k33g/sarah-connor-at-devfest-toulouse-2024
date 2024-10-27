package main

import (
	"05-more-data-rag/rag"
	"05-more-data-rag/ui"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/ollama/ollama/api"
)

func main() {

	brain := ""
	if len(os.Args) < 2 {
		// default brain
		brain = "v6"
	} else {
		brain = os.Args[1]
	}

	ctx := context.Background()
	errEnv := godotenv.Load(fmt.Sprintf("./data/brain-%s/.env", brain))

	ollamaUrl := os.Getenv("OLLAMA_HOST")
	model := os.Getenv("LLM")
	embeddingModel := os.Getenv("EMBEDDING_LLM")

	ui.Println("#ffc0c5", "🌍", ollamaUrl, "📕", model)

	client, errCli := api.ClientFromEnvironment()

	systemInstructionsFile, errInst := os.ReadFile(fmt.Sprintf("./data/brain-%s/instructions.md", brain))
	systemInstructions := string(systemInstructionsFile)

	personalityFile, errPers := os.ReadFile(fmt.Sprintf("./data/brain-%s/personality.md", brain))
	personality := string(personalityFile)

	// 🖐️🤖 all data about the terminators
	//terminatorsDataFile, errTerm := os.ReadFile(fmt.Sprintf("./data/brain-%s/terminators.data.md", brain))
	//terminatorsData := string(terminatorsDataFile)

	// Configuration
	configFile, errConf := os.ReadFile(fmt.Sprintf("./data/brain-%s/options.json", brain))
	var config map[string]interface{}
	errJsonConf := json.Unmarshal(configFile, &config)

	errorsList := errors.Join(errEnv, errCli, errInst, errPers, errConf, errJsonConf)
	if errorsList != nil {
		log.Fatal("😡:", errorsList)
	}

	// RAG section
	store := rag.MemoryVectorStore{
		Records: make(map[string]rag.VectorRecord),
	}

	// Unmarshal the store from a JSON file if it exists
	storeFile := "store.json"
	if _, err := os.Stat(storeFile); err == nil {
		file, err := os.ReadFile(storeFile)
		if err != nil {
			log.Fatal("Failed to read store file:", err)
		}
		if err := json.Unmarshal(file, &store); err != nil {
			log.Fatal("Failed to unmarshal store:", err)
		}
	}

	// Preload the model and data before starting the conversation
	//tools.PreloadModel(ctx, ollamaUrl, client, config, model, systemInstructions, personality)

	ui.Println("#FFFF00", "📝 config:", config)

	memory := []api.Message{}

	for {

		question, _ := ui.Input("#008000", fmt.Sprintf("🤖 [%s] 🧠 (%s) ask me something> ", model, brain))

		if question == "bye" {
			break
		}

		// Embbeding the question - search for the closest chunk(s)
		reqEmbedding := &api.EmbeddingRequest{
			Model:  embeddingModel,
			Prompt: question,
		}
		resp, errEmb := client.Embeddings(ctx, reqEmbedding)
		if errEmb != nil {
			fmt.Println("😡:", errEmb)
		}
		embeddingFromQuestion := rag.VectorRecord{
			Prompt:    question,
			Embedding: resp.Embedding,
		}
		similarities, _ := store.SearchTopNSimilarities(embeddingFromQuestion, 0.5, 2)

		//fmt.Println("👋 similarities:", len(similarities))

		documentsContent := "<context>\n"
		for _, similarity := range similarities {
			documentsContent += fmt.Sprintf("<doc>%s</doc>\n", similarity.Prompt)
		}
		documentsContent += "</context>"

		ui.Println("#FFFF00", "📝 similarities:", documentsContent)

		// Prompt construction
		messages := []api.Message{
			{Role: "system", Content: systemInstructions},
			{Role: "system", Content: personality},
			{Role: "system", Content: "TERMINATORS_DATA:\n" + documentsContent},
			//{Role: "user", Content: question},
		}
		// Add memory
		messages = append(messages, memory...)
		// Add the new question
		messages = append(messages, api.Message{Role: "user", Content: question})

		req := &api.ChatRequest{
			Model:    model,
			Messages: messages,
			Options:  config,
		}
		
		// Start the counter goroutine
		done := make(chan struct{})
		go func() {
			counter := 0
			for {
				select {
				case <-done:
					return
				default:
					counter++
					fmt.Printf("\r⏳ Computing... %d seconds", counter)
					time.Sleep(1 * time.Second)
				}
			}
		}()
		
		answer := ""

		respFunc := func(resp api.ChatResponse) error {
			if answer == "" {
				fmt.Println(" ✅")
				fmt.Println()
				close(done)
			}
			fmt.Print(resp.Message.Content)
			answer += resp.Message.Content
			return nil
		}

		err := client.Chat(ctx, req, respFunc)


		// Save the conversation in memory
		memory = append(
			memory,
			api.Message{Role: "user", Content: question},
			api.Message{Role: "assistant", Content: answer},
		)

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
		- what is the T-1001?
		- can you describe the T-1000 Series
	*/
}
