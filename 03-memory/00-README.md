# Garder la conversation en mémoire

<img src="imgs/sarah-connor-memory-3.jpg" width="50%" height="50%">

```golang
// Prompt construction
messages := []api.Message{
    {Role: "system", Content: systemInstructions},
    {Role: "system", Content: personality},
}
// Add memory
messages = append(messages, memory...)
// Add the new question
messages = append(messages, api.Message{Role: "user", Content: question})

// ...

// Save the conversation in memory
memory = append(
    memory,
    api.Message{Role: "user", Content: question},
    api.Message{Role: "assistant", Content: answer},
)
```

___
[👀 montre moi du code ▶️](./main.go#L48)



