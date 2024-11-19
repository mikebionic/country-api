include .env

dev:
	@go run cmd/country/main.go

db:
	@echo "Initializing countryApi database..."
	@PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d postgres \
		-f ./schemas/0.0.1_drop_db.sql
	@PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME) \
    	-f ./schemas/0.0.2_create_country_city.sql
	@echo "Has been successfully created"

build:
	@echo "Building the app, please wait..."
	@go build -o ./bin/countryApi cmd/country/main.go
	@echo "Done."