.PHONY: postgres createdb dropdb disconnectdb migrateup migratedown migrateredo migratedrop resetdb sqlc

# Start a Docker container running PostgreSQL
postgres:
	docker run --name jetcs16 -p 5432:5432 -e POSTGRES_USER=jetcs -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

# Create the journal_trends database
createdb:
	docker exec -it jetcs16 createdb -U jetcs --owner=jetcs journal_trends

# Drop the journal_trends database
dropdb:
	docker exec -it jetcs16 dropdb -U jetcs journal_trends

# Disconnect any active users before dropping the DB
disconnectdb:
	docker exec -it jetcs16 psql -U jetcs -d postgres -c "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = 'journal_trends';"

# Apply all up migrations
migrateup:
	migrate -path db/migration -database "postgres://jetcs:secret@localhost:5432/journal_trends?sslmode=disable" -verbose up

# Roll back the last migration
migratedown:
	migrate -path db/migration -database "postgres://jetcs:secret@localhost:5432/journal_trends?sslmode=disable" -verbose down

# Roll back and re-apply the last migration
migrateredo:
	make migratedown
	make migrateup

# Drop the DB safely (disconnect + drop)
migratedrop: disconnectdb dropdb

# Full reset: drop → create → migrate → regenerate Go code
resetdb: migratedrop createdb migrateup

# Generate Go code from SQL queries
sqlc:
	sqlc generate
