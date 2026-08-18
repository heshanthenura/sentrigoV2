dev:
	cd engine && go build -o bin/engine ./cmd/engine
	sudo ./engine/bin/engine

web-dev:
	cd frontend && npm run dev -- --host
