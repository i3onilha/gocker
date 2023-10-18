dev: up bash

up:
	@docker compose up -d app-dev

ps:
	@docker compose ps

logs:
	@docker compose logs app-dev --follow

build:
	@docker compose down && docker compose build --no-cache app-dev && docker compose up -d app-dev && docker compose exec app-dev bash

down:
	@docker compose down

bash:
	@docker compose exec app-dev bash

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
	@mysql -h mysql-dev -u default -p < databases/sql/mysql/schema/labels.sql

dbdown:
	@mysql -h mysql-dev -u default -p < databases/sql/mysql/down.sql

dbdump:
	@mysqldump -h mysql-dev -u root -p dbdev > databases/sql/mysql/backup/$$(date +"%Y-%m-%-d").sql

server:
	@go build -o test/ internal/cmd/server/server.go && cp env-example .env && ./test/server

server-test:
	@go build -o test/ internal/cmd/server/server.go && cp env-test .env && ./test/server
