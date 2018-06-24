build:
	go get -v -t -d ./...
	env GOOS=linux go build -ldflags="-s -w" -o bin/goal-checker goal-checker/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/goal-notifier goal-notifier/main.go
