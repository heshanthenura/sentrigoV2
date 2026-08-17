dev:
	cd engine && go build -o bin/engine ./cmd/engine
	sudo ./engine/bin/engine
