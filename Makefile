run:
	cd frontend && npm install && npm run build
	go run main.go

dev-fe:
	cd frontend && npm install && npm run dev

dev-be:
	go run main.go

build-fe:
	cd frontend && npm install && npm run build

install:
	cd frontend && npm install