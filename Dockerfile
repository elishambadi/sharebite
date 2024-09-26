# Step 1: Build the Go application in a temporary container
FROM golang:1.23-alpine AS build

# Set environment variables
ENV GO111MODULE=on
ENV GOPATH=/go
ENV DB_USER=postgres
ENV DB_NAME=sharebite
ENV DB_HOST=localhost
ENV GIN_MODE=debug
ENV SEED_DB=false

# Create a directory for the app
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the application code
COPY . .

# Build the Go app
RUN go build -o /app/sharebite-api .

# Step 2: Create a minimal image to run the app
FROM alpine:latest

# Set work directory
WORKDIR /root/

# Copy the Go app from the build container
COPY --from=build /app/sharebite-api .

# Expose the necessary port (8080 for example)
EXPOSE 8080

# Command to run the executable
CMD ["./sharebite-api"]
