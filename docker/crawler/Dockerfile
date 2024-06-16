# Build stage
FROM golang:1.22-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o crawler ./cmd/crawler/main.go

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/crawler .
COPY ./cmd/crawler/config ./cmd/crawler/config

EXPOSE 8084
CMD [ "/app/crawler" ]
