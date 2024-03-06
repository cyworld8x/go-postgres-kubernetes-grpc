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
test: 
	go test -v -cover ./..
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test

