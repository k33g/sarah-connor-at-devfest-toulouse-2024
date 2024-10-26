
Ce diagramme illustre les relations entre les principaux personnages de **Terminator 1**. Ils montrent comment Sarah Connor est au centre de l'intrigue, étant à la fois la cible du Terminator et la future mère de John Connor. Kyle Reese est envoyé pour la protéger, tandis que le Terminator est envoyé par Skynet pour l'éliminer.

```mermaid
graph TD
    A[Sarah Connor] --> B[John Connor<br>futur]
    C[Kyle Reese] --> |Protège| A
    D[Terminator T-800] --> |Cible| A
    B --> |Envoie| C
    E[Skynet] --> |Envoie| D
    A --> |Mère| B
```