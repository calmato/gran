.PHONY: setup
setup:
	cp ./.env.sample ./.env
	docker-compose build
	docker-compose run client yarn
	docker-compose run storybook yarn

.PHONY: start
start:
	docker-compose up
