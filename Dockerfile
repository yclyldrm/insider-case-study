FROM golang:1.23-alpine

WORKDIR /app

# Install required build tools and SQLite
RUN apk add --no-cache \
    gcc \
    musl-dev \
    sqlite \
    sqlite-dev

# Install swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Generate swagger docs
RUN swag init -g cmd/main.go

# Build the application with CGO enabled
RUN go build -o /app/main ./cmd/main.go

# Create a new directory for the binary
RUN mkdir -p /usr/local/bin/ && \
    cp /app/main /usr/local/bin/

# Set the binary as executable
RUN chmod +x /usr/local/bin/main

# Expose the application port
EXPOSE 9005

# Run the application
CMD ["/usr/local/bin/main"]