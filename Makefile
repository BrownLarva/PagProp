run: build 
	@./bin/app

build:
	@go build -o bin/app .	

css:
	./tailwindcss -i views/css/app.css -o public/styles.css --watch

seed:
	@go run cmd/seed/main.go

db-status:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://postgres:foobarbaz@localhost:5432/propiedades?sslmode=disable" goose -dir=db/migrations status

up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://postgres:foobarbaz@localhost:5432/propiedades?sslmode=disable" goose -dir=db/migrations up

down:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://postgres:foobarbaz@localhost:5432/propiedades?sslmode=disable" goose -dir=db/migrations down



