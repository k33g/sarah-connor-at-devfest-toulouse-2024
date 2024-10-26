#!/bin/bash
curl http://bob.local:11434/api/chat -d '{
  "model": "qwen2.5:0.5b",
  "messages": []
}'