postgres:
	docker run --name postgres -p 20241:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:16-alpine
stoppostgres:
	docker rm postgres
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres socialdb
dropdb:
	docker exec -it postgres dropdb --username=postgres socialdb	
migrateup:
	migrate -path misc/db/postgres/migration -database "postgresql://postgres:postgres@localhost:20241/socialdb?sslmode=disable" -verbose up
migratedown:
	migrate -path misc/db/postgres/migration -database "postgresql://postgres:postgres@localhost:20241/socialdb?sslmode=disable" -verbose down
sqlc:
	sqlc generate
mock:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen	
	mockgen -source=usecase/user/interface.go -destination=usecase/user/mock/user.go -package=mock
	mockgen -source=internal/user/usecases/users/interfaces.go -destination=internal/user/usecases/users/mock/service.go -package=mock
	mockgen -source=internal/user/domain/interfaces.go -destination=internal/user/domain/mock/userrepository.go -package=mock
	mockgen -package mockdb -destination=db/mock/store.go github.com/cyworld8x/go-postgres-kubernetes-grpc/db/sqlc Store

test: 
	go test -v -cover ./...  ./.... -coverprofile cover.out ./coverage/..
	go tool cover -html cover.out -o ./coverage/cover.html
proto:
	protoc --go_out=. \
    --go-grpc_out=. \
    pkg/pb/proto/*.proto
server:
	go run main.go

rebuild-db: dropdb createdb migrateup

rebuild: sqlc proto 

run-ticket:
	go run cmd/ticket/main.go
run-event-api:
	go run cmd/event/main.go
run-user:
	go run cmd/user/main.go
.PHONY: postgres stoppostgres createdb dropdb migrateup migratedown sqlc test server mock proto migrateupuserdb rebuild-db rebuild run-ticket run-user

