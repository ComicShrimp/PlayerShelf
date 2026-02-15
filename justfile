set dotenv-load := true

hot-reload:
  air

start-db:
  docker compose up -d

stop-db:
  docker compose down

get-deps:
  go mod download

update-deps:
  go get -u ./...
  go mod tidy

remove-unused-deps:
  go mod tidy

install-dev-go-deps:
  go install github.com/air-verse/air@latest
  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

create-migration migration_name:
    migrate create -ext=sql -dir=internal/infra/database/migrations {{migration_name}}

run-migrations:
    migrate -path=internal/infra/database/migrations -database $DATABASE_URL -verbose up

undo-migrations:
    migrate -path=internal/infra/database/migrations -database $DATABASE_URL -verbose down
