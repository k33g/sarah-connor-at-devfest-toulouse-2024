package flock

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"
)

type AgentFunction func() interface{}

// InstructionFunc represents a function that takes context variables and returns a string
type InstructionFunc func(map[string]interface{}) string

// Agent represents an AI agent with its configuration
type Agent struct {
	Name  string `json:"name"`
	Model string `json:"model"`
	//OllamaUrl    string `json:"ollama_url"`
	OllamaClient api.Client
	Instructions interface{}     `json:"instructions"` // Can be string or function returning string
	Functions    []AgentFunction `json:"functions"`
	//ToolChoice   *string         `json:"tool_choice,omitempty"`
	//ParallelToolCalls bool            `json:"parallel_tool_calls"`
	Options map[string]interface{}
}

// SetInstructions sets the instructions for the agent
func (a *Agent) SetInstructions(instructions interface{}) error {
	switch v := instructions.(type) {
	case string:
		a.Instructions = v
	case func() string:
		// Convert simple function to InstructionFunc
		a.Instructions = InstructionFunc(func(map[string]interface{}) string {
			return v()
		})
	case func(map[string]interface{}) string:
		a.Instructions = InstructionFunc(v)
	default:
		return fmt.Errorf("invalid instruction type: must be string, func() string, or func(map[string]interface{}) string")
	}
	return nil
}

// GetInstructions returns the current instructions as a string, using the provided context variables
func (a *Agent) GetInstructions(contextVars map[string]interface{}) string {
	switch v := a.Instructions.(type) {
	case string:
		return v
	case InstructionFunc:
		return v(contextVars)
	case func(map[string]interface{}) string:
		return v(contextVars)
	default:
		return "Invalid instruction type"
	}
}

// Response represents the response from running an agent
type Response struct {
	Messages         []api.Message          `json:"messages"`
	Agent            Agent                  `json:"agent,omitempty"`
	ContextVariables map[string]interface{} `json:"context_variables"`
}

// Orchestrator represents an API client for running agents
type Orchestrator struct {
	Ctx context.Context
}

// Run executes the agent with the given messages and context variables
func (c *Orchestrator) Run(agent Agent, messages []api.Message, contextVars map[string]interface{}, display func(string)) (Response, error) {
	// Get instructions with context variables
	instructions := agent.GetInstructions(contextVars)
	agentMessages := []api.Message{}

	agentMessages = append(agentMessages, api.Message{
		Role:    "system",
		Content: instructions,
	})
	agentMessages = append(agentMessages, messages...)

	queryChat := &api.ChatRequest{
		Model:    agent.Model,
		Messages: agentMessages,
		Options:  agent.Options,
		//Stream:   false, // TODO: make this configurable
		//Format:  "json",
	}

	var answer api.ChatResponse
	var messageContent = ""
	respFunc := func(resp api.ChatResponse) error {
		answer = resp
		messageContent += resp.Message.Content
		display(resp.Message.Content)
		//fmt.Print("🤖", resp.Message.Role, ">" , resp.Message.Content)
		return nil
	}

	err := agent.OllamaClient.Chat(c.Ctx, queryChat, respFunc)
	if err != nil {
		return Response{}, err
	}

	response := Response{
		Messages:         agentMessages,
		Agent:            agent,
		ContextVariables: contextVars,
	}

	// Add a simple response message (for demonstration)
	response.Messages = append(response.Messages, api.Message{
		//Role:    "assistant",
		Role:    answer.Message.Role,
		Content: messageContent,
	})

	return response, nil
}

func (c *Orchestrator) RunStream(agent Agent, messages []api.Message, contextVars map[string]interface{}) (Response, error) {
	return Response{}, nil
}
