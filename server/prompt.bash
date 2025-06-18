#!/bin/bash

lite="gemini-2.5-flash-lite-preview-06-17"
flash="gemini-2.5-flash"

read -p "Enter your prompt: " prompt
read -p "Choose a model (flash or lite): " model

case "$model" in
  lite)
    selected_model="$lite"
    ;;
  flash)
    selected_model="$flash"
    ;;
  *)
    echo "Error: Unknown model '$model'. Please enter 'flash' or 'lite'."
    exit 1
    ;;
esac

# send the request
curl -X POST http://localhost:3000/api/v1/chat \
  -H "Content-Type: application/json" \
  -d '{
    "prompt": "'"$prompt"'",
    "model": "'"$selected_model"'"
  }'

