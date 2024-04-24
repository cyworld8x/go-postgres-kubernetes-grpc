postgres:
	docker run --name postgres -p 20241:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:16-alpine
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres socialdb
dropdb:
	docker exec -it postgres dropdb --username=postgres socialdb	
migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:20241/socialdb?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:20241/socialdb?sslmode=disable" -verbose down
sqlc:
	sqlc generate

mock:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen	
	mockgen -source=usecase/user/interface.go -destination=usecase/user/mock/user.go -package=mock
	mockgen -package mockdb -destination=db/mock/store.go github.com/cyworld8x/go-postgres-kubernetes-grpc/db/sqlc Store

test: 
	go test -v -cover ./... -coverprofile cover.out ./coverage/..
	go tool cover -html cover.out -o ./coverage/cover.html
proto:
	protoc --go_out=. \
    --go-grpc_out=. \
    pkg/pb/proto/*.proto
server:
	go run main.go
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock proto

