
run:
	go run ./cmd/

test:
	go test ./...

.PHONY: build
build:
	rm -rf ./build/
	mkdir -p ./build/
	go build -o ./build/sel ./cmd/

hub_update:
	make build
	@hub_ctrl ${HUB_MODE} ln "$(realpath ./build/sel)" "${HOME}/.local/bin/sel"
