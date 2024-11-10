package main

import (
	"06-one-more-thing/flock"
	"06-one-more-thing/ui"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/ollama/ollama/api"
)

const (
	Red     string = "#FF0000"
	Green   string = "#00FF00"
	Blue    string = "#0000FF"
	Yellow  string = "#FFFF00"
	Orange  string = "#FFA500"
	Purple  string = "#800080"
	Pink    string = "#FFC0CB"
	Brown   string = "#A52A2A"
	Black   string = "#000000"
	White   string = "#FFFFFF"
	Gray    string = "#808080"
	Cyan    string = "#00FFFF"
	Magenta string = "#FF00FF"
)

func GetSarahInformation() (api.Client, string, string, map[string]interface{}, error) {

	ollamaUrl := os.Getenv("SARAH_OLLAMA_HOST")
	model := os.Getenv("SARAH_LLM")

	ui.Println(Cyan, "🌍", ollamaUrl, "🧠", model)

	systemInstructionsFile, errInst := os.ReadFile("./data/brain-sarah/instructions.md")
	systemInstructions := string(systemInstructionsFile)

	// Configuration
	configFile, errConf := os.ReadFile("./data/brain-sarah/settings.json")
	var config map[string]interface{}
	errJsonConf := json.Unmarshal(configFile, &config)

	url, errURL := url.Parse(ollamaUrl)
	client := api.NewClient(url, http.DefaultClient)

	errorsList := errors.Join(errInst, errConf, errJsonConf, errURL)
	if errorsList != nil {
		log.Fatal("😡:", errorsList)
	}

	return *client, model, systemInstructions , config, nil
}

func GetTerminatorInformation() (api.Client, string, string, map[string]interface{}, error) {

	ollamaUrl := os.Getenv("TERMINATOR_OLLAMA_HOST")
	model := os.Getenv("TERMINATOR_LLM")

	ui.Println(Cyan, "🌍", ollamaUrl, "🧠", model)

	systemInstructionsFile, errInst := os.ReadFile("./data/brain-terminator/instructions.md")
	systemInstructions := string(systemInstructionsFile)

	// Configuration
	configFile, errConf := os.ReadFile("./data/brain-terminator/settings.json")
	var config map[string]interface{}
	errJsonConf := json.Unmarshal(configFile, &config)

	url, errURL := url.Parse(ollamaUrl)
	client := api.NewClient(url, http.DefaultClient)

	errorsList := errors.Join(errInst, errConf, errJsonConf, errURL)
	if errorsList != nil {
		log.Fatal("😡:", errorsList)
	}

	return *client, model, systemInstructions , config, nil
}

func DisplayCounter() (chan struct{}, string) {
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
	result := ""
	return done, result
}

func main() {
	errEnv := godotenv.Load("./.env")
	if errEnv != nil {
		log.Fatal("😡:", errEnv)
	}

	ctx := context.Background()

	flockOrchestrator := flock.Orchestrator{
		Ctx: ctx,
	}

	client, model, systemInstructions, config, err := GetSarahInformation()
	if err != nil {
		log.Fatal(err)
	}

	sarah := flock.Agent{
		Name:         "Sarah",
		Model:        model,
		OllamaClient: client,
		Options:      config,
	}
	ui.Println(Green, sarah.OllamaClient)

	sarah.SetInstructions(systemInstructions)


	client, model, systemInstructions, config, err = GetTerminatorInformation()
	if err != nil {
		log.Fatal(err)
	}
	terminator := flock.Agent{
		Name:         "Terminator",
		Model:        model,
		OllamaClient: client,
		Options:      config,
	}
	ui.Println(Green, terminator.OllamaClient)

	terminator.SetInstructions(systemInstructions)

	ui.Println(Green, "\n=== Starting conversation between Sarah and The Terminator at the coffee shop ===\n")
	// Initialize conversation with a starter
	conversation := []api.Message{
		{Role: "system", Content: "This is a conversation between Sarah and Model 101. They should actively engage with each other, responding to questions and comments."},
		{Role: "user", Content: "Sarah and Model 101 meet at a coffee shop. Start a conversation about weekend plans."},
	}

	var sarahResponse flock.Response
	var terminatorResponse flock.Response

	for i := 0; i < 3; i++ {

		// Sarah's turn
		ui.Println(Orange, "👩 Sarah Connor> ")
		sarahCounter, result := DisplayCounter()


		sarahResponse, err = flockOrchestrator.Run(sarah, conversation, map[string]interface{}{},
			func(s string) {
				if result == "" {
					fmt.Println(" ✅")
					close(sarahCounter)
				}
				result += s
				ui.Print(Orange, s)
			},
		)
		if err != nil {
			fmt.Println("😡 bob error:", err)
		}
		lastSarahMessage := sarahResponse.Messages[len(sarahResponse.Messages)-1]
		//ui.Println(Green, "🐼 Bob: "+lastSarahMessage.Content)

		// Add Sarah's message to conversation history
		conversation = append(conversation, api.Message{
			Role:    "user",
			Content: lastSarahMessage.Content,
		})

		fmt.Println()
		fmt.Println()

		// Terminator's turn
		ui.Println(Red, "🤖 Model 101> ")
		terminatorCounter, result := DisplayCounter()

		terminatorResponse, err = flockOrchestrator.Run(terminator, conversation, map[string]interface{}{},
			func(s string) {
				if result == "" {
					fmt.Println(" ✅")
					close(terminatorCounter)
				}
				result += s
				ui.Print(Red, s)
			},
		)
		if err != nil {
			fmt.Println("😡 sam error:", err)
		}
		lastTerminatorMessage := terminatorResponse.Messages[len(terminatorResponse.Messages)-1]
		//ui.Println(Red, "🐻 Sam: "+lastTerminatorMessage.Content)

		// Add Sam's message to conversation history
		conversation = append(conversation, api.Message{
			Role:    "user",
			Content: lastTerminatorMessage.Content,
		})

		fmt.Println()
		fmt.Println()

	}

	ui.Println(Green, "\n=== End of conversation ===\n")

}
