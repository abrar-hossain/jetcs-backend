.PHONY: postgres createdb dropdb migrateup migratedown sqlc

postgres:
	docker run --name jetcs16 -p 5432:5432 -e POSTGRES_USER=jetcs -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it jetcs16 createdb -U jetcs --owner=jetcs journal_trends

dropdb:
	docker exec -it jetcs16 dropdb -U jetcs journal_trends

migrateup:
	migrate -path db/migration -database "postgres://jetcs:secret@localhost:5432/journal_trends?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://jetcs:secret@localhost:5432/journal_trends?sslmode=disable" -verbose down

sqlc:
	sqlc generate
