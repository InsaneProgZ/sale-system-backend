# Stage 1: Download dependencies
FROM golang:alpine AS dependencies

WORKDIR /build

# Copy Go module files
COPY /app/go.mod /app/go.sum ./

# Download dependencies
RUN go mod download

# Stage 2: Build the application
FROM dependencies AS builder

# Copy the rest of the application code
RUN pwd && ls -l

COPY /app .

# Build the Go binary
RUN GOARCH=amd64 GOOS=linux go build -o main

# Stage 3: Final image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /build/main ./

EXPOSE 8080

# Use the binary as entrypoint
ENTRYPOINT ["./main"]
