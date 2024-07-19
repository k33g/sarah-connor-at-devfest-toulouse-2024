#!/bin/bash

OLLAMA_URL="http://localhost:11434"
MODEL="qwen:0.5b"

# Instruction
read -r -d '' SYSTEM_CONTENT <<- EOM

EOM

# Question
SYSTEM_CONTENT=$(echo ${SYSTEM_CONTENT} | tr -d '\n')

read -r -d '' USER_CONTENT <<- EOM

EOM

USER_CONTENT=$(echo ${USER_CONTENT} | tr -d '\n')


