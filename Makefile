install:
	go get -v -t -d ./...

build: install
	env GOOS=linux go build -ldflags="-s -w" -o bin/goal-checker goal-checker/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/goal-notifier goal-notifier/main.go

test: install
	go test -v ./...

clean:
	rm -rf bin/
	go clean

.PHONY: install gbuild test clean
