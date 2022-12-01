code: upddev vim

dev: upddev bashdev

vim:
	@docker compose exec app-dev vim

up:
	@docker compose up

upddev:
	@docker compose up -d app-dev

ps:
	@docker compose ps

logs:
	@docker compose logs --follow

build-dev:
	@docker-compose down && docker-compose build --no-cache app-dev && docker-compose up -d app-dev && docker-compose exec app-dev bash

down:
	@docker compose down

bashdev:
	@docker compose exec app-dev bash
