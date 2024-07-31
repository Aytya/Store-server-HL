.PHONY: build run

build:
	docker-compose build

down:
	docker-compose down

run:
	docker-compose up