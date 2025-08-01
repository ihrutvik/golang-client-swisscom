# Golang Parallel HTTP Client

This is a simple Golang program that sends parallel POST requests to a backend service endpoint.

## ✅ Features

- Sends multiple parallel POST requests to a backend service
- Each request contains dynamic JSON data with unique service, resource, and owner details
- Number of requests per batch and number of batches is configurable

## 🔧 Prerequisites

- Go should be installed (`go version` to verify)
- Backend service should be running on `http://localhost:8080`

## 🚀 How to Run

1. Open terminal and go to the project folder.
2. Make the script executable:
```bash
chmod +x run.sh

```

3. Run the script:
```bash
./run.sh
```

Or run manually:
```bash
go run main.go --url=http://localhost:8080/api/v1/services --parallel=10 --steps=5
```

You can change `--parallel` and `--steps` values to control the load.
