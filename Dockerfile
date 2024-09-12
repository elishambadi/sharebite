# Use the official Golang image from the Docker Hub
FROM golang:1.20

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Install sqlc and goose
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Build the Go app
RUN go build -o main ./cmd/server

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]
