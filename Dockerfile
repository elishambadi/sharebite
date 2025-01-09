# Step 1: Build the Go application in a temporary container
FROM golang:1.23-alpine AS build

# Set environment variables for Go build
ENV GO111MODULE=on
ENV GOPATH=/go
ENV DB_USER=postgres
ENV DB_NAME=sharebite
ENV DB_HOST=localhost
ENV GIN_MODE=debug
ENV SEED_DB=false

# Create a directory for the Go app
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Install Go dependencies
RUN go mod download

# Copy the application code
COPY . .

# Build the Go app
RUN go build -o /app/sharebite-api .

# Step 2: Build the Node.js assets (CSS, JavaScript, etc.)
FROM node:16-alpine AS node-build

# Set the working directory for Node.js build
WORKDIR /app

# Copy the application code including package.json and package-lock.json
COPY . .

# Install dependencies (including TailwindCSS) and build the assets
RUN npm install -D tailwindcss \
    && npm run build:css # This runs your build script to generate the final assets


RUN ls -l /app/assets/styles > /app/files-list-new.txt

# Step 3: Create the minimal runtime image
FROM alpine:latest as runtime

# Install Node.js, npm, and necessary dependencies in the final image
RUN apk add --no-cache nodejs npm

# Set the working directory for the runtime container
WORKDIR /root/

# Copy the Go binary from the build container
COPY --from=build /app/sharebite-api .
RUN mkdir -p assets
COPY --from=node-build /app/assets/ ./assets/
COPY --from=node-build /app/files-list-new.txt .
COPY --from=build /app/.env .env

# Expose the necessary port (8080 for example)
EXPOSE 8081

# Command to run the Go executable
CMD ["./sharebite-api"]
