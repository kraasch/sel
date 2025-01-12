
run:
	go run ./cmd/

test:
	go test ./...

.PHONY: build
build:
	rm -rf ./build/
	mkdir -p ./build/
	go build -o ./build/sel ./cmd/
