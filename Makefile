.PHONY: all build run logs generate migrate seed stop

# Variables
DOCKER_COMPOSE_FILE=docker-compose.yaml

all: build

build:
	docker-compose -f $(DOCKER_COMPOSE_FILE) build --no-cache
run:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d
stop:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down
logs:
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f
generate:
	go generate ./src/...
migrate:
	docker-compose -f $(DOCKER_COMPOSE_FILE) run --rm graphql go run src/cmd/migrate/main.go --action=migrate
migrate/rollback:
	docker-compose -f $(DOCKER_COMPOSE_FILE) run --rm graphql go run src/cmd/migrate/main.go --action=rollback
migrate/refresh:
	docker-compose -f $(DOCKER_COMPOSE_FILE) run --rm graphql go run src/cmd/migrate/main.go --action=refresh
seed:
	docker-compose -f $(DOCKER_COMPOSE_FILE) run --rm graphql go run src/cmd/seed/main.go