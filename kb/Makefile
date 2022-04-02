.PHONY: all
all:
	mkdir -p ./bin
	go mod tidy
	go build -o ./bin/kb ./cmd

.PHONY: test
test:
	go test ./...

.PHONY: gen
gen:
	go run ./gen.go
	go fmt ./temp

.PHONY: clean
clean:
	rm -rf ./bin
