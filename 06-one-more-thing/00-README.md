# One more thing: les Agents

### Sarah et le Terminator partent en week-end

<img src="imgs/sarah-en-week-end.jpg" width="50%" height="50%">


#### Petite expérimentation inspirée par Swarm d'OpenAI

Le code se compose de deux parties principales :

1. `flock`: Une bibliothèque qui définit:
- Un `Agent` qui représente une IA avec un nom, un modèle, des instructions et des options
- Un `Orchestrator` qui gère l'exécution des agents et leurs conversations
- Des fonctionnalités pour gérer les instructions et les réponses

2. `main`: Un programme qui utilise `flock` pour:
- Créer deux agents: **Sarah** et un Terminator (**Model 101**)
- Configurer chaque agent avec ses propres **instructions** et modèle LLM
- Orchestrer une conversation entre eux dans un café
- Afficher la conversation avec un compteur de temps pour chaque réponse



```mermaid
flowchart TB
    subgraph main
        env[".env File"]
        sarah_conf["Sarah Config<br/>(instructions.md + settings.json)"]
        term_conf["Terminator Config<br/>(instructions.md + settings.json)"]
    end

    subgraph flock["Flock Library"]
        Agent
        Orchestrator
        Response
    end

    subgraph runtime["Runtime"]
        conv["Conversation Loop"]
        display["Display Counter"]
    end

    env --> main
    sarah_conf --> Agent
    term_conf --> Agent
    
    Agent --> Orchestrator
    Orchestrator --> Response
    
    Response --> conv
    conv --> display
    display --> conv

    subgraph ollama["Ollama API"]
        api["API Client"]
    end

    Agent --> api
    api --> Response

```

Points clés du fonctionnement :

1. Configuration:
- Les agents sont configurés via des fichiers (.env, instructions.md, settings.json)
- Chaque agent a son propre modèle LLM et ses propres instructions

2. Conversation:
- Se déroule en tours alternés entre Sarah et Terminator
- Chaque message est ajouté à l'historique de la conversation
- Un compteur visuel (⏳) indique le temps de réflexion

3. Communication:
- Utilise l'API Ollama pour communiquer avec les modèles LLM
- Les réponses sont streamées (affichées progressivement)
- Les messages sont formatés avec des couleurs différentes pour chaque agent

**Le système est conçu de manière modulaire, ce qui permet de créer facilement différents types d'agents avec différentes personnalités et instructions.**

```mermaid
classDiagram
    class Agent {
        +String Name
        +String Model
        +Client OllamaClient
        +Interface Instructions
        +AgentFunction[] Functions
        +Map Options
        +SetInstructions(interface) error
        +GetInstructions(Map) String
    }

    class Orchestrator {
        +Context Ctx
        +Run(Agent, Message[], Map, func) Response
        +RunStream(Agent, Message[], Map) Response
    }

    class Response {
        +Message[] Messages
        +Agent Agent
        +Map ContextVariables
    }

    class AgentFunction {
        <<interface>>
        +interface Execute()
    }

    class InstructionFunc {
        <<interface>>
        +String Execute(Map)
    }

    Agent --> AgentFunction
    Agent --> InstructionFunc
    Orchestrator --> Agent
    Orchestrator --> Response
    Response --> Agent
```