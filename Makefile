.PHONY: setup
setup:
	cp ./.env.sample ./.env
	docker-compose build
	docker-compose run client yarn
	docker-compose run storybook yarn

.PHONY: start
start:
	docker-compose up

.PHONY: client
client.setup:
	cd ./client && yarn

client.lint:
	cd ./client && yarn lint

client.build:
	cd ./client && yarn build
