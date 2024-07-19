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

	"github.com/parakeet-nest/parakeet/completion"
	"github.com/parakeet-nest/parakeet/llm"
)

func main() {
	ollamaUrl := os.Getenv("OLLAMA_URL")
	// if working from a container
	//ollamaUrl := "http://host.docker.internal:11434"

	model := os.Getenv("MODEL")
	//model := "tinyllama"

	systemContent, err := os.ReadFile("system.md")
	if err != nil {
		log.Fatal("😡:", err)
	}

	contextContent, err := os.ReadFile("context.md")
	if err != nil {
		log.Fatal("😡:", err)
	}

	userContent, err := os.ReadFile("user.md")
	if err != nil {
		log.Fatal("😡:", err)
	}

	fmt.Println(ollamaUrl)
	fmt.Println(model)
	fmt.Println(userContent)

	options := llm.Options{
		Temperature:   0.2,
		RepeatLastN:   2,
		RepeatPenalty: 2.0,
		TopK: 10,
		TopP: 0.5,
	}

	/*
				Top_p (Nucleus Sampling): Limits the response to the top-p percentage of probability mass,
				ensuring the model picks from the most likely words, making the response more focused.

				  "options": {
					"temperature": 0.1,
					"repeat_penalty": 2,
					"repeat_last_n": 2,
					"top_k": 10,
					"top_p": 0.5
				  }

		> - `top_k`: Reduces the probability of generating nonsense. A higher value (e.g. 100) will give more diverse answers, while a lower value (e.g. 10) will be more conservative. (Default: 40)
		> - `top_p`: Works together with top-k. A higher value (e.g., 0.95) will lead to more diverse text, while a lower value (e.g., 0.5) will generate more focused and conservative text. (Default: 0.9)
		> - `repeat_last_n`: Sets how far back for the model to look back to prevent repetition. (Default: 64, 0 = disabled, -1 = num_ctx)
		> - `repeat_penalty`: Sets how strongly to penalize repetitions. A higher value (e.g., 1.5) will penalize repetitions more strongly, while a lower value (e.g., 0.9) will be more lenient. (Default: 1.1)
	*/

	query := llm.Query{
		Model: model,
		Messages: []llm.Message{
			{Role: "system", Content: string(systemContent)},
			{Role: "system", Content: string(contextContent)},
			{Role: "user", Content: string(userContent)},
		},
		Options: options,
	}

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
