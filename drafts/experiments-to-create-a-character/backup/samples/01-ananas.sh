#!/bin/bash

OLLAMA_URL="http://localhost:11434"
MODEL="pizzaiolo"

read -r -d '' USER_CONTENT <<- EOM
Donne moi une recette de pizza avec de l'ananas et du miel
EOM

USER_CONTENT=$(echo ${USER_CONTENT} | tr -d '\n')

read -r -d '' DATA <<- EOM
{
  "model":"${MODEL}",
  "options": {
    "temperature": 0.1,
    "repeat_penalty": 2,
    "repeat_last_n": 2,
    "top_k": 10,
    "top_p": 0.5
  },
  "prompt": "${USER_CONTENT}",
  "stream": false
}
EOM

curl ${OLLAMA_URL}/api/generate \
   -H "Content-Type: application/json" \
   -d "${DATA}" | jq -c '{ response }'

