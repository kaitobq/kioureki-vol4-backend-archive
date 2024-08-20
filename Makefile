.PHONY: up down
COMPOSE=docker-compose
DOCKER=docker
SQL_CONTAINER_NAME=kioureki-db
up:
	@${COMPOSE} up -d

dev: 
	@${COMPOSE} up

down:
	@${COMPOSE} down

exec:
	@${DOCKER} exec -it ${SQL_CONTAINER_NAME} /bin/bash

