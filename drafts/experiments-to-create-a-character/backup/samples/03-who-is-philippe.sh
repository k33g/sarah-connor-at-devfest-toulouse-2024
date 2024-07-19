#!/bin/bash

OLLAMA_URL="http://localhost:11434"
MODEL="riviera"

read -r -d '' USER_CONTENT <<- EOM
Who is Philippe Charrière?
EOM

USER_CONTENT=$(echo ${USER_CONTENT} | tr -d '\n')

read -r -d '' DATA <<- EOM
{
  "model":"${MODEL}",
  "options": {
  },
  "prompt": "${USER_CONTENT}",
  "stream": false
}
EOM

curl ${OLLAMA_URL}/api/generate \
   -H "Content-Type: application/json" \
   -d "${DATA}" | jq -c '{ response }'



