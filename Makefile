all: check copyenv build up

# check docker and docker-compose are available
check:
ifeq (,$(shell which docker))
	$(error "command docker not found")
endif
	docker -v
ifeq (,$(shell which docker-compose))
	$(error "command docker-compose not found")
endif
	docker-compose -v

copyenv:
ifeq (,$(wildcard ./.env))
	# no .env file, copy .env.example to .env
	cp ./.env.example ./.env
endif
	cat ./.env

build:
	docker-compose build

up:
	docker-compose up