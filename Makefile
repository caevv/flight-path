SHELL=/bin/bash

.PHONY: build
build:
	CGO_ENABLED=0 go build -a -o ./build/bin/svc cmd/main.go

.PHONY: runbuild
runbuild:
	./build/bin/svc

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: gorun
gorun:
	go run cmd/main.go

.PHONY: lint
lint:
	golangci-lint run ./... --fix
