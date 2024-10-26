`[La naissance de Sarah Connor]`
## Principe du prompt

```mermaid
graph TD
    A[Prompt: messages] -->|Premier élément| B[Message 1]
    A -->|Deuxième élément| C[Message 2]
    
    B --> D["Role: 'system'"]
    B --> E["Content: systemInstructions"]
    
    C --> F["Role: 'user'"]
    C --> G["Content: question"]

    H["[]api.Message{...}"]
    H -->|Définit| A
    
    style H fill:#f9f,stroke:#333,stroke-width:2px
```

```golang
messages := []api.Message{
    {Role: "system", Content: systemInstructions},
    {Role: "user", Content: question},
}
```
___
[◀️ Previous](./README.md#la-naissance-de-sarah-connor) | [Next: Requête ▶️](./02-request.md#principe-de-la-requête)
