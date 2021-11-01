include .env

## build: Build Dockerfile.
build:
	@echo "Building Go Binary..."
	docker build --no-cache -t goweather_web .

## start-server: Start in development mode. Gets reloaded automatically when code changes.
start-server:
	@echo "Running Server..."
	docker-compose up -d --build --force-recreate

## stop-server: Stop development mode.
stop-server:
	@echo "Stopping Server..."
	docker-compose down && docker system prune

## watch-logs: Display logs in the console.
watch-logs:
	docker-compose logs -f

## clean: Remove all unused locale Volumes and remove all stopped containers.
clean:
	docker system prune -f
	docker volume prune -f

## web-container: exec web-container.
web-container:
	docker exec -itu root go-app /bin/bash

## db-container: exec db container.
db-container:
	docker exec -it go-db psql psql -U${PGUSER} -h${APP_HOST} -d${PGDATABASE}

## migrate-create: create new migration. You must pass the name of the table as an argument to table_name.Example "make migrate-create table_name=users"
migrate-create:
	migrate create -ext sql -dir db/migrations -seq create_$(table_name)_table

## migrate-up: execute migration.If you want to specify the step of migrations, pass the step parameter. Example "make migrate-up step=2"
migrate-up:
	migrate -path ./db/migrations/ -database ${DATABASE_URL}'?sslmode=disable' up $(step)

## migrate-down: down migration. If you want to specify the step of migrations, pass the step parameter. Example "make migrate-up step=2"
migrate-down:
	migrate -path ./db/migrations/ -database ${DATABASE_URL}'?sslmode=disable' down $(step)

help: Makefile
	@echo
	@echo "Available commands:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo