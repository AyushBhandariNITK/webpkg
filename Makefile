build:
	go build -o bin/webpkg main.go

run:
	./bin/webpkg

test:
	go test -v ./...


