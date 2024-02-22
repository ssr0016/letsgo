# Variables
DATABASE_URL = "postgresql://greenlight:pa55word@localhost:5432/postgres?sslmode=disable"
MIGRATIONS_PATH = migrations

# Target to apply migrations
migrate-up:
	migrate -database $(DATABASE_URL) -path $(MIGRATIONS_PATH) -verbose up

# Target to revert migrations
migrate-down:
	migrate -database $(DATABASE_URL) -path $(MIGRATIONS_PATH) -verbose down

.PHONY: migrate-up migrate-down


# go get -u github.com/golang-migrate/migrate/v4/cmd/migrate
# migrate -database "postgresql://postgres:secret@localhost:5432/postgres?sslmode=disable" -path db/migrations up
#sqlc init