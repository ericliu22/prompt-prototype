#!/bin/bash

read -p "Enter in a propmt" prompt

curl -X POST http://localhost:3000/api/v1/chat \
  -H "Content-Type: application/json" \
  -d '{
    "prompt": "'"$prompt"'",
    "model": "gemini-2.5-flash-lite"
  }'
