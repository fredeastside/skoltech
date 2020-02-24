#source .env

build:
	docker-compose up -d --build

up:
	docker-compose up -d

start: up

stop:
	docker-compose stop

down:
	docker-compose down

logs:
	docker-compose logs -f
