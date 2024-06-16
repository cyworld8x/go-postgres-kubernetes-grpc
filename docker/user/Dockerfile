# Build stage
FROM golang:1.22-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/user/main.go
RUN apk add --no-cache curl
RUN  curl -L https://github.com/golang-migrate/migrate/releases/download/v4.12.2/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
# COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY ./cmd/user/config ./cmd/user/config
COPY start.sh .
RUN chmod +x start.sh
COPY wait-for.sh .
RUN chmod +x  wait-for.sh
COPY misc/db/postgres/migration ./migration

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
