/*
Topic: Parakeet
Generate a chat completion with Ollama and parakeet
The output is streamed
*/

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/parakeet-nest/parakeet/completion"
	"github.com/parakeet-nest/parakeet/embeddings"
	"github.com/parakeet-nest/parakeet/llm"
)

func main() {
	ollamaUrl := "http://localhost:11434"
	// if working from a container
	//ollamaUrl := "http://host.docker.internal:11434"

	model := "qwen2:0.5b" //🔴

	var embeddingsModel = "all-minilm" // This model is for the embeddings of the documents
	store := embeddings.MemoryVectorStore{
		Records: make(map[string]llm.VectorRecord),
	}

	//	data, err := os.ReadFile("context-lyon.txt")
	data, err := os.ReadFile("context-all.txt")
	if err != nil {
		log.Fatal("😡:", err)
	}

	var docs = strings.Split(string(data), "------")

	// Create embeddings from documents and save them in the store
	for idx, doc := range docs {
		fmt.Println("Creating embedding from document ", idx)
		embedding, err := embeddings.CreateEmbedding(
			ollamaUrl,
			llm.Query4Embedding{
				Model:  embeddingsModel,
				Prompt: doc,
			},
			strconv.Itoa(idx),
		)
		if err != nil {
			fmt.Println("😡:", err)
		} else {
			store.Save(embedding)
		}
	}

	systemContent := `Your name is Bob and you are a French Gin expert.
	Use only the context below to answer the user question.`

	userContent := `What gin are made in Lyon?`
	//userContent := `What gin are made in Paris and Versaille?`

	// Create an embedding from the question
	embeddingFromQuestion, err := embeddings.CreateEmbedding(
		ollamaUrl,
		llm.Query4Embedding{
			Model:  embeddingsModel,
			Prompt: userContent,
		},
		"question",
	)
	if err != nil {
		log.Fatalln("😡:", err)
	}
	fmt.Println("🔎 searching for similarity...")

	similarities, _ := store.SearchSimilarities(embeddingFromQuestion, 0.65) 
	// if ~ 1 the best similarities
	// with 0.3, a lot of similarities
	// with 0.6, less similarities -> wh have reduced the context

	fmt.Println("🎉 similarities:", len(similarities))

	documentsContent := embeddings.GenerateContextFromSimilarities(similarities)

	fmt.Println("📝:", documentsContent)

	options := llm.Options{
		Temperature:   0.3,
		RepeatLastN:   2,
		RepeatPenalty: 2.0,
		TopP: 0.9,
	}

	/*
		Top_p (Nucleus Sampling): Limits the response to the top-p percentage of probability mass, 
		ensuring the model picks from the most likely words, making the response more focused.
	*/



	query := llm.Query{
		Model: model,
		Messages: []llm.Message{
			{Role: "system", Content: systemContent},
			{Role: "system", Content: documentsContent},
			{Role: "user", Content: userContent},
		},
		Options: options,
	}

	fmt.Println()
	fmt.Println("===========================================================")
	fmt.Println()

	_, err = completion.ChatStream(ollamaUrl, query,
		func(answer llm.Answer) error {
			fmt.Print(answer.Message.Content)
			return nil
		})

	fmt.Println()
	
	if err != nil {
		log.Fatal("😡:", err)
	}

}
