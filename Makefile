## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## run: run the project
.PHONY: run
run:
	@go run ./cmd/app

.PHONY: migrate
## migrate: apply migrations stored
migrate:
	@go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations

.PHONY: migrate-test
## migrate-test: apply test migrations
migrate-test:
	@go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./tests/migrations --migrations-table=migrations_test

.PHONY: test
## test: run tests
test: 
	@go test -v ./tests