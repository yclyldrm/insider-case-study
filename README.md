# Insider Case Study

A service that handles asynchronous message sending with Redis caching and webhook integration.

## Features

- Asynchronous message sending via webhooks
- Redis caching for message tracking
- Job management for message processing
- RESTful API endpoints
- Docker containerization

## Prerequisites

- Go 1.23+
- Docker and Docker Compose
- Redis

## Installation & Setup

1. Clone the repository:
```
git clone https://github.com/yclyldrm/insider-case-study.git
```

2. Install dependencies
```
go mod download
```

3. Create an `.env` file in the root directory and add values like examples
```
DB_NAME=db_name.db
WEBHOOK_URL=webhook_url
PORT=3001
````

4. Run the application
```
docker-compose up -d --build
```
`or` 
```
go run cmd/main.go
```


The server will start on port 9005.

## 📚 API Documentation

### Endpoints

- `GET /api/messages` - Get all sent messages
- `GET /api/messages/{id}` - Get a specific message by ID
- `POST /api/job-status` - Change job status (enable/disable)

### Swagger Documentation
Access the Swagger documentation at:
```
http://localhost:9005/swagger/index.html
```

## 🏗 Project Structure

```
├── cmd/
│   └── main.go
├── config/
│   └── config.go
├── internal/
│   ├── api/
│   ├── application/
│   ├── domain/
│   ├── infrastructure/
│   └── interfaces/
├── pkg/
│   └── job.go
└── .gitignore
└── Dockerfile
└── docker-compose.yml
└── go.mod
└── go.sum
└── Licence
└── README.md
```

## 🔄 Message Processing Flow

1. Messages are stored in SQLite
2. Job service picks unprocessed messages every 2 minutes
3. Messages are sent through webhook
4. Message status is updated in database
5. Message details are cached in Redis

## ⚙️ Configuration Options

- Message size limit: 100 characters
- Job interval: 2 minutes
- Server timeout: 15 seconds
- Batch processing limit: 2 messages per cycle


Project Link: https://github.com/yclyldrm/insider-case-study
