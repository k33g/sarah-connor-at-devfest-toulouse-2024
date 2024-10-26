# La personnalité de Sarah Connor
<img src="imgs/sarah-connor-gun.jpg" width="50%" height="50%">

Donner une personnalité à Sarah Connor est un élément clé pour la rendre crédible et cohérente. 

### ⏺ Ajouter plus de contexte au prompt

```golang
messages := []api.Message{
    {Role: "system", Content: systemInstructions},
    {Role: "system", Content: personality},
    {Role: "user", Content: question},
}
```

### ⏺ Changer les paramètres de la requête

#### 4 versions du 🧠 cerveau de Sarah Connor

___
[👀 montre moi du code ▶️](./main.go#L16)


