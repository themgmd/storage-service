include .env

build:
	go build -o .bin/main cmd/main.go
run: build
	./.bin/main
dev:
	go run cmd/main.go

migrate_up:
	migrate -path ./schema -database postgres://admin:admonql@localhost:5432/storage?sslmode=disable up

migrate_down:
	migrate -path ./schema -database postgres://admin:admonql@localhost:5432/storage?sslmode=disable down