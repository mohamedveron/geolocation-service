
VERSION = $(shell printf "%s.%s" \
		$$(git rev-list --count HEAD) \
		$$(git rev-parse --short HEAD))

generate:
	@echo :: getting generator
	go get -v -d
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen

	@echo :: generating code

	oapi-codegen -package api -generate chi-server,types api/api.yml > api/api.gen.go

test: generate
	@echo :: run tests
	go test -v ./tests

run: generate
	@echo :: start the http server
	go run main.go

build:  $(OUTPUT)
	CGO_ENABLED=0 GOOS=linux go build -o bin/app \
		-ldflags "-X main.version=$(VERSION)" \
		-gcflags "-trimpath $(GOPATH)/src"

all: generate build tests run


$(OUTPUT):
	mkdir -p $(OUTPUT)