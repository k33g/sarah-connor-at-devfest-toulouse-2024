package main

import (
	"05-more-data-rag/rag"
	"05-more-data-rag/txt"
	"05-more-data-rag/ui"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/ollama/ollama/api"
)

var (
	brain = flag.String("brain", "v6", "")
)

func TestGenerateChunk(t *testing.T) {

	ctx := context.Background()
	errEnv := godotenv.Load(fmt.Sprintf("./data/brain-%s/.env", *brain))
	embeddingModel := os.Getenv("EMBEDDING_LLM")

	client, errCli := api.ClientFromEnvironment()

	// 🖐️🤖 all data about the terminators
	terminatorsDataFile, errTerm := os.ReadFile(fmt.Sprintf("./data/brain-%s/terminators.data.md", *brain))
	terminatorsData := string(terminatorsDataFile)

	errorsList := errors.Join(errEnv, errCli, errTerm)
	if errorsList != nil {
		log.Fatal("😡:", errorsList)
	}

	// RAG section
	store := rag.MemoryVectorStore{
		Records: make(map[string]rag.VectorRecord),
	}
	// Split the data
	chunks := txt.SplitTextWithDelimiter(terminatorsData, "<!-- SPLIT -->")

	for index, chunk := range chunks {
		ui.Println("#FFFF00", "chunk nb", index, ":")
		ui.Println("#FF06FF", chunk)

		req := &api.EmbeddingRequest{
			Model:  embeddingModel,
			Prompt: chunk,
		}
		resp, errEmb := client.Embeddings(ctx, req)
		if errEmb != nil {
			fmt.Println("😡:", errEmb)
		}
		//ui.Println("#FF061C", "📦", resp.Embedding)

		record, errSave := store.Save(rag.VectorRecord{
			Prompt:    chunk,
			Embedding: resp.Embedding,
		})
		if errSave != nil {
			fmt.Println("😡:", errSave)
		}
		ui.Println("#FF061C", "📦", record.Embedding)

	}

	// Marshal the store to JSON
	storeJSON, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		log.Fatal("Failed to marshal store to JSON:", err)
	}

	// Write the JSON to a file
	storeFile := "store.json"
	err = os.WriteFile(storeFile, storeJSON, 0644)
	if err != nil {
		log.Fatal("Failed to write store to file:", err)
	}

	fmt.Println("✅ Store persisted to", storeFile)

}
