#!/bin/bash
# from your project root
set -a
source .env
set +a
go run ./cmd/main.go

