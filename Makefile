include .env

build:
	go build -o .bin/app.exe ./cmd

run:
	go run cmd -race

init:
	cp .env.example .env

# Сборка проекта
up:
	docker compose up -d --build

migrate-up:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:5436/${DB_DATABASE}?sslmode=disable up

migrate-down:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:5436/${DB_DATABASE}?sslmode=disable down -all

# Создание таблицы
create-table:
	migrate create -ext sql -dir ./migrations -seq create_$(table_name)_table

# Создание таблицы
rebuild:
	docker-compose up -d --no-deps --build $(container)

swag-init:
	swag init -g cmd/main.go