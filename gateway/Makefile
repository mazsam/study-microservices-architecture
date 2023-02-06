.PHONY: validate clean generate build run

all: generate build

validate:
	swagger validate ./api/swagger.yml

clean:
	rm -rf gateway-server
	go clean -i .
	
build: clean
	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix cgo ./cmd/gateway-server

run: 
	./gateway-server --port=8080 --host=0.0.0.0

run-local:
	go run cmd/gateway-server/main.go --port=8080

generate: validate
	go generate ./...
