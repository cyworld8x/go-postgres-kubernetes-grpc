# --- Build stage ---
FROM golang:1.24-alpine3.21 AS builder

WORKDIR /app

# Only copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary statically for a minimal final image
RUN CGO_ENABLED=0 GOOS=linux go build -o user ./cmd/user/main.go

# --- Run stage ---
FROM alpine:3.21

WORKDIR /app

# Only copy the built binary and necessary config
COPY --from=builder /app/user .
COPY ./cmd/user/config ./cmd/user/config

# Ensure the binary has execution permissions (if needed)
RUN chmod +x /app/user

# Expose the port
EXPOSE 5004

# Set the entrypoint
ENTRYPOINT ["/app/user"]
