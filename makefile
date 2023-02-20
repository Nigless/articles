include .env
export

migrate-up:
	go run ./migrate/main.go up
.PHONY: migrate

run:
	air -build.cmd "go build -o tmp/main cmd/main.go"