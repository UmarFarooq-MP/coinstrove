# Base image for building the Go app
FROM golang:1.21 AS go-builder

# Set the working directory inside the container
WORKDIR /go/app

# Copy go.mod and go.sum files to cache dependencies
COPY go.mod go.sum ./

# Download dependencies, leveraging Docker cache
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go app
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/app/main ./cmd/main.go

# Use a minimal image for the final application
FROM alpine:latest AS base

# Install necessary runtime dependencies (optional, depends on your app needs)
RUN apk add --no-cache curl wget

# Copy the built Go app from the builder stage
COPY --from=go-builder /go/app/main /main

# Set the default command to run the Go app
CMD ["/main"]
