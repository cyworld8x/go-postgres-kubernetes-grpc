# Build stage
FROM golang:1.22-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o event ./cmd/event/main.go

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/event .
COPY ./cmd/event/config ./cmd/event/config

EXPOSE 8083
CMD [ "/app/event" ]
