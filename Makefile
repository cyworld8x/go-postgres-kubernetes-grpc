postgres:
	docker run --rm --name postgres -p 20241:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:16-alpine
redis:
	docker run --rm --name redis -p 6379:6379 -d redis:7.2.5-alpine --save 60 1 --loglevel warning
rabbitmq:
	docker run --rm --name rabbitmq -p 5677:5672 -p 15677:15672 -d -e RABBITMQ_DEFAULT_USER=rabbitmq -e RABBITMQ_DEFAULT_PASS=rabbitmq rabbitmq:3-management
asynq-management:
	docker run --rm \
    --name asynqmon \
    -p 8080:8080 \
    hibiken/asynqmon:latest --redis-addr=host.docker.internal:6379\
dynamodb:
	docker run --rm --name dynamodb -p 8000:8000 -d amazon/dynamodb-local
stoppostgres:
	docker rm postgres
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres socialdb
dropdb:
	docker exec -it postgres dropdb --username=postgres socialdb	
swagger:
	docker run --rm -v $(pwd):/code ghcr.io/swaggo/swag:latest
gen-swagger:
	swag init -d internal/event/application/api/,internal/event/domain/  -g server.go  -o internal/event/application/api/swagger/docs/
	swag init -d internal/crawler/application/api/,internal/crawler/domain/  -g server.go  -o internal/crawler/application/api/swagger/docs/
migrateup:
	migrate -path misc/db/postgres/migration -database "postgresql://postgres:postgres@localhost:20241/socialdb?sslmode=disable" -verbose up
migratedown:
	migrate -path misc/db/postgres/migration -database "postgresql://postgres:postgres@localhost:20241/socialdb?sslmode=disable" -verbose down
sqlc:
	sqlc generate
mock:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen	
	mockgen -source=internal/user/usecases/users/interfaces.go -destination=internal/user/usecases/users/mock/service.go -package=mock
	mockgen -source=internal/user/domain/interfaces.go -destination=internal/user/domain/mock/repository.go -package=mock
	mockgen -source=internal/user/usecases/sessions/interfaces.go -destination=internal/user/usecases/sessions/mock/service.go -package=mock
	mockgen -source=internal/user/usecases/users/interfaces.go -destination=internal/user/usecases/users/mock/service.go -package=mock
	mockgen -source=internal/user/infrastructure/repository/postgres -destination=internal/user/infrastructure/repository/mock -package=mock
	
test: 
	go test -v -cover ./...  ./.... -coverprofile cover.out ./coverage/..
	go tool cover -html cover.out -o ./coverage/cover.html
proto:
	protoc --go_out=. \
    --go-grpc_out=. \
    pkg/pb/proto/*.proto

install-playwright:
	go run github.com/playwright-community/playwright-go/cmd/playwright@latest install --with-deps

rebuild-db: stoppostgres postgres createdb migrateup

rebuild: sqlc proto 

run-ticket:
	go run cmd/ticket/main.go
run-event-api:
	go run cmd/event/main.go
run-crawler-api:
	go run cmd/crawler/main.go
run-user:
	go run cmd/user/main.go
.PHONY: postgres stoppostgres createdb dropdb migrateup migratedown sqlc test server mock proto rebuild-db rebuild run-ticket run-user gen-swagger install-playwright redis asynq-management

