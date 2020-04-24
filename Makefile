CLIENT_DIR := ./client
STORYBOOK_DIR := ./story
USER_API_DIR := ./api/user

##################################################
# Container Commands - All
##################################################
.PHONY: setup
setup:
	cp ./.env.sample ./.env
	docker-compose build
	$(MAKE) install

.PHONY: install
install:
	docker-compose run --rm client yarn
	docker-compose run --rm storybook yarn
	docker-compose run --rm group_api make setup
	docker-compose run --rm todo_api make setup
	docker-compose run --rm user_api make setup

.PHONY: start
start:
	docker-compose up

##################################################
# Container Commands - Only API
##################################################
.PHONY: api-setup
api-setup:
	docker-compose -f docker-compose.api.yml build
	$(MAKE) api-install

.PHONY: api-install
api-install:
	docker-compose -f docker-compose.api.yml run group_api make setup
	docker-compose -f docker-compose.api.yml run todo_api make setup
	docker-compose -f docker-compose.api.yml run user_api make setup

.PHONY: api-start
api-start:
	docker-compose -f docker-compose.api.yml up

##################################################
# Container Commands - Terraform
##################################################
.PHONY: terraform-stg-setup
terraform-stg-setup:
	docker-compose run --rm terraform make init ENV=stg

.PHONY: terraform-stg-lint
terraform-stg-lint:
	docker-compose run --rm terraform make fmt ENV=stg

.PHONY: terraform-stg-plan
terraform-stg-plan:
	docker-compose run --rm terraform make plan ENV=stg

.PHONY: terraform-stg-apply
terraform-stg-apply:
	docker-compose run --rm terraform make apply ENV=stg

.PHONY: terraform-stg-destroy
terraform-stg-destroy:
	docker-compose run --rm terraform make destroy ENV=stg

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

.PHONY: client-build
client-build:
	cd ${CLIENT_DIR}; yarn build

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

.PHONY: storybook-build
storybook-build:
	cd ${STORYBOOK_DIR}; yarn build-storybook

##################################################
# Local Commands - API (ToDo Service)
##################################################
.PHONY: todo-api-setup
todo-api-setup:
	cp ${TODO_API_DIR}/.envrc.sample ${TODO_API_DIR}/.envrc
	cd ${TODO_API_DIR}; make setup

.PHONY: todo-api-start
todo-api-start:
	cd ${TODO_API_DIR}; make run

.PHONY: todo-api-lint
todo-api-lint:
	cd ${TODO_API_DIR}; make fmt
	cd ${TODO_API_DIR}; make lint

.PHONY: todo-api-test
todo-api-test:
	cd ${TODO_API_DIR}; make test

.PHONY: todo-api-build
todo-api-build:
	cd ${TODO_API_DIR}; make build

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

.PHONY: user-api-build
user-api-build:
	cd ${USER_API_DIR}; make build
