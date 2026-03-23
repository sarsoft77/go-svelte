.PHONY: run dev build clean install

BINARY_NAME=myapp

run: dev

dev:
	@echo "Building frontend..."
	@cd frontend && npm run build >/dev/null 2>&1
	@echo "Starting backend (Go)..."
	@cd backend && go run . &
	@sleep 3
	@echo "Starting frontend (Vite)..."
	@cd frontend && npm run dev
	@echo ""
	@echo "Frontend: http://localhost:5173"
	@echo "Backend API: http://localhost:8080"

backend:
	cd backend && go run .

frontend:
	cd frontend && npm run dev

build:
	cd frontend && npm run build
	cd backend && go build -o ../$(BINARY_NAME) .

install:
	cd frontend && npm install

clean:
	rm -f $(BINARY_NAME)
	rm -rf backend/data
	cd frontend && rm -rf dist

help:
	@echo "Available targets:"
	@echo "  make run         - Run both backend and frontend"
	@echo "  make dev         - Same as run"
	@echo "  make backend     - Run only Go backend"
	@echo "  make frontend    - Run only frontend dev server"
	@echo "  make build       - Build both applications"
	@echo "  make install    - Install frontend dependencies"
	@echo "  make clean       - Clean build artifacts"
	@echo "  make help        - Show this help"
