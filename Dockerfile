# Stage 1: Download dependencies
FROM golang:alpine3.20 AS dependencies

WORKDIR /build

# Copy Go module files
COPY /src/go.mod /src/go.sum ./

# Download dependencies
RUN go mod download

# Stage 2: Build the application
FROM dependencies AS builder

# Copy the rest of the application code
RUN pwd && ls -l

COPY /src .

# Build the Go binary
RUN GOARCH=amd64 GOOS=linux go build -o main

# Stage 3: Final image
FROM alpine:2.7

WORKDIR /src

# Copy the binary from the builder stage
COPY --from=builder /build/main ./

EXPOSE 8080

# Use the binary as entrypoint
ENTRYPOINT ["./main"]
