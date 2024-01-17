include .env
export $(shell sed 's/=.*//' .env)

GOLANGCI_LINT_URL=https://raw.githubusercontent.com/golangci/golangci-lint
GOLANGCI_LINT_VERSION=v1.55.2
HTMX_URL=https://unpkg.com/htmx.org
HTMX_VERSION=1.9.10

prepare: tools build-frontend templ deps

database-dev-up: postgresql-local-up migrate-up models

database-dev-down: postgresql-local-down

tools:
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/cosmtrek/air@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/lqs/sqlingo/sqlingo-gen-postgres@latest
	wget -O /dev/stdout ${GOLANGCI_LINT_URL}/${GOLANGCI_LINT_VERSION}/install.sh | sh -s ${GOLANGCI_LINT_VERSION}
	npm install -D tailwindcss@latest
	npm install flowbite

build-frontend:
	mkdir -p ./static/css ./static/js
	npx tailwindcss -i ./assets/css/app.css -o ./static/css/app.min.css --minify
	wget ${HTMX_URL}@${HTMX_VERSION}/dist/htmx.min.js -O ./static/js/htmx.min.js -o /dev/null
	wget ${HTMX_URL}@${HTMX_VERSION}/dist/ext/json-enc.js -O ./static/js/json-enc.js -o /dev/null

check:
	golangci-lint run
	go test -v -race ./...

templ:
	templ generate

deps:
	go mod tidy


postgresql-local-up:
	docker-compose --env-file .env up -d db
	@sleep 2

postgresql-local-down:
	docker-compose --env-file .env down

migrate-up:
	goose -dir migrations postgres "host=${DATABASE_HOST} user=${DATABASE_USER} password=${DATABASE_PASSWORD} dbname=${DATABASE_NAME} sslmode=disable" up

migrate-down:
	goose -dir migrations postgres "host=${DATABASE_HOST} user=${DATABASE_USER} password=${DATABASE_PASSWORD} dbname=${DATABASE_NAME} sslmode=disable" down

models:
	sqlingo-gen-postgres "host=${DATABASE_HOST} user=${DATABASE_USER} password=${DATABASE_PASSWORD} dbname=${DATABASE_NAME} sslmode=disable" > domain/model/model.go

run:
	air

build: prepare
	go build -o ./build/bootstrap .

image:
	DOCKER_BUILDKIT=1 docker build -f docker/app/Dockerfile -t kuchy .

run-image: image
	docker run -p 0.0.0.0:3000:3000/tcp kuchy
