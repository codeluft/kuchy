include .env
export $(shell sed 's/=.*//' .env)

GOLANGCI_LINT_URL=https://raw.githubusercontent.com/golangci/golangci-lint
GOLANGCI_LINT_VERSION=v1.55.2
HTMX_URL=https://unpkg.com/htmx.org
HTMX_VERSION=1.9.10

.PHONY: prepare
prepare: tools build-frontend templ sql deps

.PHONY: database-dev-up
database-dev-up: postgresql-local-up migrate-up

.PHONY: database-dev-down
database-dev-down: postgresql-local-down

.PHONY: tools
tools:
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/cosmtrek/air@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	wget -O /dev/stdout ${GOLANGCI_LINT_URL}/${GOLANGCI_LINT_VERSION}/install.sh | sh -s ${GOLANGCI_LINT_VERSION}
	npm install

.PHONY: build-frontend
build-frontend:
	mkdir -p ./static/css ./static/js
	npx tailwindcss -i ./assets/css/app.css -o ./static/css/app.min.css --minify
	wget ${HTMX_URL}@${HTMX_VERSION}/dist/htmx.min.js -O ./static/js/htmx.min.js -o /dev/null
	wget ${HTMX_URL}@${HTMX_VERSION}/dist/ext/json-enc.js -O ./static/js/json-enc.js -o /dev/null

.PHONY: check
check:
	./bin/golangci-lint run
	go test -v -race ./...

.PHONY: templ
templ:
	templ generate

.PHONY: deps
deps:
	go mod tidy


.PHONY: postgresql-local-up
postgresql-local-up:
	docker-compose --env-file .env up -d db
	@sleep 2

.PHONY: postgresql-local-down
postgresql-local-down:
	docker-compose --env-file .env down

.PHONY: migrate-up
migrate-up:
	goose -dir database/migrations postgres "host=${DATABASE_HOST} user=${DATABASE_USER} password=${DATABASE_PASSWORD} dbname=${DATABASE_NAME} sslmode=disable" up

.PHONY: migrate-down
migrate-down:
	goose -dir database/migrations postgres "host=${DATABASE_HOST} user=${DATABASE_USER} password=${DATABASE_PASSWORD} dbname=${DATABASE_NAME} sslmode=disable" down

.PHONY: sql
sql:
	sqlc generate

.PHONY: run
run:
	air

.PHONY: build
build: prepare
	go build -o ./build/bootstrap .

.PHONY: image
image:
	DOCKER_BUILDKIT=1 docker build -f docker/app/Dockerfile -t kuchy .

.PHONY: run-image
run-image: image
	docker run -p 0.0.0.0:3000:3000/tcp kuchy
