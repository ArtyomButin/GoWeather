## build: Build Dockerfile.
build:
	@echo "Building Go Binary..."
	docker build --no-cache -f Dockerfile -t main .

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

web-container:
	docker exec -itu root go-app /bin/bash

help: Makefile
	@echo
	@echo "Available commands:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo