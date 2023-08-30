include .env

build:
	go build -o .bin/app.exe ./cmd

run:
	go run cmd -race

init:
	cp .env.example .env

up:
	docker compose up -d --build

migrate-up:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:5436/${DB_DATABASE}?sslmode=disable up

migrate-down:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://root:root@localhost:5436/db_pet?sslmode=disable down -all