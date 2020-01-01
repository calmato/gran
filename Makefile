CLIENT_DIR := ./client
STORYBOOK_DIR := ./story
USER_API_DIR := ./api/user

##################################################
# Container Commands
##################################################
.PHONY: setup
setup:
	cp ./.env.sample ./.env
	docker-compose build
	$(MAKE) install

.PHONY: install
install:
	docker-compose run client yarn
	docker-compose run storybook yarn
	docker-compose run user_api make setup

.PHONY: start
start:
	docker-compose up

##################################################
# Local Commands - Client
##################################################
.PHONY: client-setup
client-setup:
	cp ${CLIENT_DIR}/.envrc.sample ${CLIENT_DIR}/.envrc
	$(MAKE) client-install

.PHONY: client-install
client-install:
	cd ${CLIENT_DIR}; yarn

.PHONY: client-start
client-start:
	cd ${CLIENT_DIR}; yarn dev

.PHONY: client-lint
client-lint:
	cd ${CLIENT_DIR}; yarn lint

.PHONY: client-test
client-test:
	cd ${CLIENT_DIR}; yarn test

##################################################
# Local Commands - Storybook
##################################################
.PHONY: storybook-setup
storybook-setup:
	cp ${STORYBOOK_DIR}/.envrc.sample ${STORYBOOK_DIR}/.envrc
	$(MAKE) storybook-install

.PHONY: storybook-install
storybook-install:
	cd ${STORYBOOK_DIR}; yarn

.PHONY: storybook-start
storybook-start:
	cd ${STORYBOOK_DIR}; yarn storybook

.PHONY: storybook-lint
sotybook-lint:
	cd ${STORYBOOK_DIR}; yarn lint

.PHONY: storybook-test
storybook-test:
	cd ${STORYBOOK_DIR}; yarn test

##################################################
# Local Commands - API (User Service)
##################################################
.PHONY: user-api-setup
user-api-setup:
	cp ${USER_API_DIR}/.envrc.sample ${USER_API_DIR}/.envrc
	cd ${USER_API_DIR}; make setup

.PHONY: user-api-start
user-api-start:
	cd ${USER_API_DIR}; make run

.PHONY: user-api-lint
user-api-lint:
	cd ${USER_API_DIR}; make fmt
	cd ${USER_API_DIR}; make lint

.PHONY: user-api-test
user-api-test:
	cd ${USER_API_DIR}; make test
