dev: up bash

up:
	@docker compose up -d sagemcom-dev

ps:
	@docker compose ps

logs:
	@docker compose logs sagemcom-dev --follow

build:
	@docker compose down && docker compose build --no-cache sagemcom-dev && docker compose up -d sagemcom-dev && docker compose exec sagemcom-dev bash

down:
	@docker compose down

stop:
	@docker compose stop

bash:
	@docker compose exec sagemcom-dev bash

bashoracle:
	@docker compose exec oracle-dev bash

bashmysql:
	@mysql -h mysql-dev -u default -p

t:
	@clear;go test ./...

tv:
	@clear;go test -v ./...

cover:
	@go test -coverprofile=test/coverage.out ./... && go tool cover -html=test/coverage.out -o test/coverage.html && go run test/cover.go

sqlc:
	@sqlc generate -f sqlc.mysql.yaml

dbup:
	@mysql -h mysql-dev -u default -p < databases/sql/mysql/schema/import_pallets_serials.sql

dbdown:
	@mysql -h mysql-dev -u default -p < databases/sql/mysql/down.sql

dbdump:
	@mysqldump -h mysql-dev -u root -p dbdev > databases/sql/mysql/backup/$$(date +"%Y-%m-%-d").sql

start:
	@cp .env-example .env && air

start-test:
	@go build -o test/ internal/cmd/server/server.go && cp .env-test .env && ./test/server
