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
	@docker-compose run client yarn

.PHONY: client-lint
client-lint:
	@docker-compose run client yarn lint

.PHONY: client-build
client-build:
	@docker-compose run client yarn build
