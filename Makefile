dev:
	go build -o bin/sentrigo ./cmd/sentrigo
	sudo ./bin/engine

build:
	go build -o bin/sentrigo ./cmd/sentrigo

web-dev:
	cd frontend && npm run dev -- --host
