.phony: all
all:
	mkdir -p ./bin
	go mod vendor
	go mod tidy
	go build -o ./bin/kb ./cmd

.phony: foo
foo:
	./bin/kb foo

.phony: test
test:
	go test ./...

.phony: clean
clean:
	rm -rf ./bin
