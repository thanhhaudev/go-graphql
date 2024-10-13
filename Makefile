.PHONY: all build run stop

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
