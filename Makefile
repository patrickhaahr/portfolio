run: build
	@./bin/app

build:
	@go build -tags dev -o bin/app .

build-heroku:
	@go build -tags heroku -o bin/app

css:
	tailwindcss -i views/css/app.css -o public/styles.css --watch

templ:
	templ generate --watch --proxy=http://localhost:3000
