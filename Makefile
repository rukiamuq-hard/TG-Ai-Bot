COMPOSE=docker compose -f deployments/docker/docker-compose.yml

.PHONY: up down logs build rebuild ps shell

up:
	$(COMPOSE) up -d --build

down:
	$(COMPOSE) down

restart:
	$(COMPOSE) down
	$(COMPOSE) up -d --build

logs:
	$(COMPOSE) logs -f

build:
	$(COMPOSE) build

rebuild:
	$(COMPOSE) down
	$(COMPOSE) build --no-cache
	$(COMPOSE) up -d

ps:
	$(COMPOSE) ps

shell:
	docker exec -it tg-ai-bot sh
