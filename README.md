# Go Message Broker

A simple message broker implemented in Go with HTTP API endpoints.

## Features
- Thread-safe message queue
- REST API endpoints for publishing and consuming messages
- JSON message format

## Usage
```bash
# Start the server
go run main.go

# Publish a message
curl -X POST http://localhost:8080/publish \
  -H "Content-Type: application/json" \
  -d '{"id": "msg1", "payload": "Hello World!"}'

# Consume a message
curl -X GET http://localhost:8080/consume
```

## API Endpoints
- `POST /publish` - Add a message to the queue
- `GET /consume` - Get and remove the next message from the queue