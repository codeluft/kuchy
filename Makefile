include .env
export $(shell sed 's/=.*//' .env)

GOLANGCI_LINT_VERSION="v1.55.2"

prepare: tools assets templ deps

database-dev-up: postgresql-local-up migrate-up models

database-dev-down: postgresql-local-down

tools:
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/cosmtrek/air@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/lqs/sqlingo/sqlingo-gen-postgres@latest
	wget -O /dev/stdout \
		https://raw.githubusercontent.com/golangci/golangci-lint/${GOLANGCI_LINT_VERSION}/install.sh | \
		sh -s ${GOLANGCI_LINT_VERSION}

assets:
	wget https://unpkg.com/htmx.org@1.9.10/dist/htmx.min.js \
		-O ./static/js/htmx.min.js -o /dev/null
	wget https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css \
		-O ./static/css/bootstrap.min.css -o /dev/null
	wget https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js \
		-O ./static/js/bootstrap.min.js -o /dev/null

check:
	golangci-lint run || true
	go test -v -race ./... || true

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
