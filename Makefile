include .env

PHONY: build
build:
	go build -o ./bin/app .

PHONY: run
run:
	PORT=${PORT} ./bin/app

PHONY: docker-clean
docker-clean:
	docker container prune
	docker image prune
	docker network prune
	docker volume prune

PHONY: updev
updev:
	GO_ENV=${GO_ENV_DEV} docker-compose up

PHONY: uptest
uptest:
	GO_ENV=${GO_ENV_TEST} docker-compose up

PHONY: upprod
upprod:
	GO_ENV=${GO_ENV_PROD} docker-compose up

PHONY: down
down:
	docker-compose down --rmi all

DEFAULT_GOAL := run