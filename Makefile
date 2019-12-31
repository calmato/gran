ENV ?= dev

.PHONY: setup
setup:
	cp ./.env.sample ./.env
	docker-compose build
	docker-compose run client yarn
	docker-compose run storybook yarn

.PHONY: start
start:
	docker-compose up

.PHONY: client-setup
client-setup:
ifeq ($(ENV), ci)
	@echo [Info] ENV is ${ENV}
	@cd ./client && yarn
else
	@echo [Info] ENV is ${ENV}
	@docker-compose run client yarn
endif

.PHONY: client-lint
client-lint:
ifeq ($(ENV), ci)
	@echo [Info] ENV is ${ENV}
	@cd ./client && yarn lint
else
	@echo [Info] ENV is ${ENV}
	@docker-compose run client yarn lint
endif

.PHONY: client-build
client-build:
ifeq ($(ENV), ci)
	@echo [Info] ENV is ${ENV}
	@cd ./client && yarn build
else
	@echo [Info] ENV is ${ENV}
	@docker-compose run client yarn build
endif
