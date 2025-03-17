include .env

PHONY: build
build:
	go build -o ./bin/app .

PHONY: run
run:
	PORT=${PORT} ./bin/app

PHONY: dbuild
dbuild:
	PORT=${PORT} docker build -t go-app:0.1.1 -f Dockerfile.multistage .

PHONY: drun
drun:
	docker run go-app:0.1.0

PHONY: docker-clean
docker-clean:
	docker container prune
	docker image prune
	docker network prune
	docker volume prune

PHONY: updev
updev:
	GO_ENV=${GO_ENV_DEV} docker-compose up --build

PHONY: uptest
uptest:
	GO_ENV=${GO_ENV_TEST} docker-compose up --build

PHONY: up
up:
	GO_ENV=${GO_ENV} docker-compose up

PHONY: down
down:
	docker-compose down --rmi all

DEFAULT_GOAL := down