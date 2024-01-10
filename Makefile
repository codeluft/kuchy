prepare: tools assets templ

tools:
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/cosmtrek/air@latest

assets:
	wget https://unpkg.com/htmx.org@1.9.10/dist/htmx.min.js \
		-O ./static/js/htmx.min.js -o /dev/null
	wget https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css \
		-O ./static/css/bootstrap.min.css -o /dev/null
	wget https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js \
		-O ./static/js/bootstrap.min.js -o /dev/null

templ:
	templ generate
	go mod tidy

run:
	air

build: setup templates
	go build -o ./build/bootstrap .

image:
	DOCKER_BUILDKIT=1 docker build -f docker/app/Dockerfile -t kuchy-frontend .

run-image: image
	docker run -p 0.0.0.0:3000:3000/tcp kuchy-frontend
