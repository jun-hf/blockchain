build:
	go build -o bin/blockchain

run: build
	./bin/blockchain