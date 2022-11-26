code: upd vim

dev: upd bash

vim:
	@docker compose exec app-dev vim

up:
	@docker compose up

upd:
	@docker compose up -d

ps:
	@docker compose ps

logs:
	@docker compose logs --follow

build:
	@docker compose up -d --build

down:
	@docker compose down

bash:
	@docker compose exec app-dev bash
