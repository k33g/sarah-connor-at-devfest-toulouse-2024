package main

import (
	"fmt"
	"log"

	"github.com/parakeet-nest/gollama"
	"github.com/parakeet-nest/gollama/enums/option"

)

func main() {
	ollamaUrl := "http://localhost:11434"
	// if working from a container
	//ollamaUrl := "http://host.docker.internal:11434"
	model := "qwen2:0.5b"

	systemContent := `You are Sarah Connor. Your job is to know everything about the terminators`

	//userContent := `Give me a list of all the terminators models`
	//userContent := `Who is Sarah Connor?`
	userContent := `Who is John Connor for you?`

	options := gollama.SetOptions(map[string]interface{}{
		option.Temperature: 0.0,
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

	// 	fullAnswer, err := gollama.ChatStream(ollamaUrl, query,

	_, err := gollama.ChatStream(ollamaUrl, query,
		func(answer gollama.Answer) error {
			fmt.Print(answer.Message.Content)
			return nil
		})

	//fmt.Println("📝 Full answer:")
	//fmt.Println(fullAnswer.Message.Role)
	//fmt.Println(fullAnswer.Message.Content)

	if err != nil {
		log.Fatal("😡:", err)
	}
}
