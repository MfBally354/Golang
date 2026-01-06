# Makefile untuk memudahkan Docker commands

.PHONY: help build run stop clean logs shell test

# Default target
help:
	@echo "Available commands:"
	@echo "  make build         - Build Docker images"
	@echo "  make run           - Run containers"
	@echo "  make stop          - Stop containers"
	@echo "  make restart       - Restart containers"
	@echo "  make clean         - Remove containers and images"
	@echo "  make logs          - Show logs"
	@echo "  make logs-api      - Show API logs"
	@echo "  make logs-backend  - Show backend logs"
	@echo "  make shell         - Access container shell"
	@echo "  make test          - Run tests"
	@echo "  make dev           - Run in development mode"

# Build Docker images
build:
	docker-compose build

# Run containers
run:
	docker-compose up -d
	@echo "✅ Containers running!"
	@echo "API: http://localhost:8090"
	@echo "Backend: http://localhost:8091"

# Run in foreground (untuk development)
dev:
	docker-compose up

# Stop containers
stop:
	docker-compose down

# Restart containers
restart: stop run

# Remove containers, networks, and images
clean:
	docker-compose down --rmi all --volumes
	@echo "✅ Cleaned up!"

# Show all logs
logs:
	docker-compose logs -f

# Show API logs only
logs-api:
	docker-compose logs -f api

# Show backend logs only
logs-backend:
	docker-compose logs -f backend

# Access container shell
shell:
	docker exec -it golang-api sh

# Run tests inside container
test:
	docker-compose exec api go test ./...

# Build specific service
build-api:
	docker-compose build api

build-backend:
	docker-compose build backend

# Status of containers
status:
	docker-compose ps

# Prune unused Docker resources
prune:
	docker system prune -af --volumes
	@echo "✅ Docker resources pruned!"
