install:
	go get -v -t -d ./...

build: install
	env GOOS=linux go build -ldflags="-s -w" -o bin/goal-checker goal-checker/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/goal-notifier goal-notifier/main.go

test: install
	go test -v -race ./...

coveralls: install
	go get github.com/mattn/goveralls
	go test -v -cover -race -coverprofile=coverage.out ./...
	goveralls -coverprofile=coverage.out -service=circle-ci -repotoken=${COVERALLS_TOKEN}

clean:
	rm -rf bin/
	go clean

.PHONY: install gbuild test clean
