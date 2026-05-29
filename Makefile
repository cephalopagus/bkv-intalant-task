include .env
export  

run:
	@go run cmd/app/main.go

env-up:
	@docker-compose up -d intalant-postgres

env-down:
	@docker compose down intalant-postgres

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder


migrate-up:
	@docker-compose run --rm -e GOOSE_COMMAND=up migrate

migrate-down:
	@docker-compose run --rm -e GOOSE_COMMAND=down migrate

migrate-status:
	@docker-compose run --rm -e GOOSE_COMMAND=status migrate

