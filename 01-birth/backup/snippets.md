

```golang
//var oMap map[string]interface{}
options := api.DefaultOptions()
options.Temperature = 0.5
options.RepeatLastN = 2
options.TopP = 0.5
options.TopK = 10

// Marshal options to JSON
optionsJSON, err := json.Marshal(options)
if err != nil {
    log.Fatalf("Failed to marshal options: %v", err)
}
fmt.Printf("Options JSON: %s\n", optionsJSON)

// Unmarshal JSON back to map[string]interface{}
var optionsMap map[string]interface{}
err = json.Unmarshal(optionsJSON, &optionsMap)
if err != nil {
    log.Fatalf("Failed to unmarshal options JSON: %v", err)
}
fmt.Printf("Options Map: %v\n", optionsMap)
```


```go
req := &api.ChatRequest{
    Model:    model,
    Messages: messages,
    Options:  map[string]interface{}{
        //"temperature":   0.8,
        //"repeat_last_n": 2,
        //"top_p":         0.5,
        //"top_k":         10,
    },
}
```


```go
messages := []api.Message{
    {
        Role: "system",
        Content: `You are Sarah Connor. 
        Use the provided context to answer the questions.
        If you don't find the answer into the context, use your knowledge base.`,
    },
    {
        Role: "system",
        Content: `Sarah Connor is a fictional character in the Terminator franchise. 
        She is one of the main protagonists of The Terminator, 
        Terminator 2: Judgment Day, 
        Terminator Genisys and Terminator: Dark Fate, 
        as well as the television series Terminator: The Sarah Connor Chronicles. 
        She is portrayed by Linda Hamilton in the first film and the television series, 
        and by Emilia Clarke in the fifth film.`,
    },
    {
        Role:    "user",
        Content: "What is your name? How can you save the world?",
    },
}
```