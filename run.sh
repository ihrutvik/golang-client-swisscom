#!/bin/bash
echo "Running Golang HTTP client..."
go run main.go --url=http://localhost:8080/api/v1/services --parallel=10 --steps=5