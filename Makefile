.PHONY: tailwind
tailwind:
	npx tailwindcss -i ./assets/styles/input.css build -o ./static/css/style.css

.PHONY: templ
templ:
	templ generate