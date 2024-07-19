#!/bin/bash

OLLAMA_URL="http://localhost:11434"
#MODEL="qwen:0.5b"
MODEL="qwen2:0.5b"

read -r -d '' USER_CONTENT <<- EOM
Où habites tu ?
EOM

USER_CONTENT=$(echo ${USER_CONTENT} | tr -d '\n')

OLLAMA_URL=$OLLAMA_URL MODEL=$MODEL USER_CONTENT=$USER_CONTENT go run main.go

