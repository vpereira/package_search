APP_NAME=package_search

.PHONY: build run clean

build:
	go build -o $(APP_NAME)

test:
	go test -v ./...

lint: 
	go vet && go fmt ./...

run:
	./$(APP_NAME)

clean:
	rm -f $(APP_NAME)

gosec:
	go install  github.com/securego/gosec/v2/cmd/gosec@latest
	gosec -exclude=G104,G114 ./...
