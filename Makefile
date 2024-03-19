
build:
	   @go build -o bin/gobank
run: build
	   @./bin/gobank
dev: build
	   nodemon --exec go run main.go --signal SIGTERM

test:
	   @go test -v ./...