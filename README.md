# Go Authentication Server

A simple authentication server built with Go that provides user registration, login, and profile management using JWT tokens and MongoDB.

## Project Structure
```
go-auth/
├─ go.mod          # Go module dependencies
├─ main.go         # Server setup and main entry point
├─ config.go       # Configuration variables
├─ db.go           # MongoDB connection and utilities
├─ model.go        # Data models
├─ handlers.go     # HTTP request handlers
├─ middleware.go   # Authentication middleware
```

## Prerequisites

- Go 1.25.3 or later
- MongoDB running locally (default: mongodb://localhost:27017)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/standardhub/go-auth.git
cd go-auth
```

2. Install dependencies:
```bash
go mod download
```

## Configuration

The server uses the following default configuration in `config.go`:

- MongoDB URI: `mongodb://localhost:27017`
- MongoDB Database: `goauth`
- JWT Token Expiry: 24 hours
- Server Port: 8080

## Running the Server

1. Make sure MongoDB is running locally
2. Start the server:
```bash
go run .
```

The server will start on `http://localhost:8080`

## API Endpoints

### Register a new user
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "yourpassword"}'
```

### Login
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "yourpassword"}'
```

### Get User Profile (Protected Route)
```bash
curl http://localhost:8080/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Development

Build the project:
```bash
go build ./...
```

## Error Responses

The API returns appropriate HTTP status codes and JSON error messages:

- 400: Bad Request (invalid input)
- 401: Unauthorized (invalid credentials or token)
- 404: Not Found
- 500: Internal Server Error
