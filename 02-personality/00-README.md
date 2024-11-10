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

Voici la traduction en français :

- `top_k` : Réduit la probabilité de générer du non-sens. Une valeur plus élevée (par exemple 100) donnera des réponses plus diverses, tandis qu'une valeur plus basse (par exemple 10) sera plus conservatrice. (Par défaut : 40)
- `top_p` : Fonctionne en conjonction avec top_k. Une valeur plus élevée (par exemple 0.95) conduira à un texte plus diversifié, tandis qu'une valeur plus basse (par exemple 0.5) générera un texte plus ciblé et conservateur. (Par défaut : 0.9)

#### 3 versions du 🧠 cerveau de Sarah Connor
#### 🖐️ >>> Passer directement à la version 3

___
[👀 montre moi du code ▶️](./main.go#L16)


