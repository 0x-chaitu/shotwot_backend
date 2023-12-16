postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=postgres --owner=postgres shotwot

dropdb:
	docker exec -it postgres13 dropdb shotwot

install_migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/shotwot?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/shotwot?sslmode=disable" -verbose down

.PHONY: createdb dropdb postgres install_migrate