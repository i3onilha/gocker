dev: up bash

up:
	@docker compose up -d app-dev

ps:
	@docker compose ps

logs:
	@docker compose logs app-dev --follow

down:
	@docker compose down

bash:
	@docker compose exec app-dev bash

t:
	@clear;go test ./...

tv:
	@clear;go test -v ./...

cover:
	@go test -coverprofile=test/coverage.out ./... && go tool cover -html=test/coverage.out -o test/coverage.html && go run test/cover.go

server:
	@go build -o test/ internal/cmd/server/server.go && cp .env-example .env && ./test/server

server-test:
	@go build -o test/ internal/cmd/server/server.go && cp .env-test .env && ./test/server
