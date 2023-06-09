dev: upddev bashdev

up:
	@docker compose up

upddev:
	@docker compose up -d app-dev

ps:
	@docker compose ps

logs:
	@docker compose logs --follow

build-dev:
	@docker compose down && docker compose build --no-cache app-dev && docker compose up -d app-dev && docker compose exec app-dev bash

build-prod:
	@docker compose down && docker compose build --no-cache app-prod && docker compose up -d app-prod

down:
	@docker compose down

bashdev:
	@docker compose exec app-dev bash

t:
	@clear;go test ./...

tv:
	@clear;go test -v ./...

cover:
	@go test -coverprofile=test/coverage.out ./... && go tool cover -html=test/coverage.out -o test/coverage.html && go run test/cover.go
