dev: up bash

up:
	@docker compose up -d

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

dbup:
	@mysql -h mysql-dev -u default -p < databases/sql/mysql/schema/create.sql

mysqlbash:
	@mysql -h mysql-dev -u default -p
