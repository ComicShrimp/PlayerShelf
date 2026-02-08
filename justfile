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
