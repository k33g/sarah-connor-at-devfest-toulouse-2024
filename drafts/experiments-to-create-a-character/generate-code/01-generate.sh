#!/bin/bash

OLLAMA_URL="http://localhost:11434"
#MODEL="qwen:0.5b"
#MODEL="qwen2:0.5b"
#MODEL="deepseek-coder:1.3b"
#MODEL="phi3:mini" // 🟠
#MODEL="codegemma:2b"
#MODEL="starcoder:3b"
#MODEL="starcoder:1b"
#MODEL="llama3:latest" // 🟠
#MODEL="mistral:latest" // 🔴
#MODEL="granite-code:3b" // 🔴
#MODEL="rouge/replete-coder-qwen2-1.5b:Q8" // 🟠
#MODEL="stablelm2:latest"
MODEL="stable-code:latest"
#MODEL=""

OLLAMA_URL=$OLLAMA_URL MODEL=$MODEL go run main.go > dungeon.js


