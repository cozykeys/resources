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

buf/download:
	mkdir -p ./bin
	wget https://github.com/bufbuild/buf/releases/download/v1.4.0/buf-Linux-x86_64
	mv ./buf-Linux-x86_64 ./bin/buf
	chmod +x ./bin/buf

buf/generate:
	./bin/buf generate
