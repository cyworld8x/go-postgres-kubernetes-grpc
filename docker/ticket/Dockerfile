# Build stage
FROM golang:1.22-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o ticket ./cmd/ticket/main.go

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/ticket .
COPY ./cmd/ticket/config ./cmd/ticket/config

EXPOSE 8082
CMD [ "/app/ticket" ]
