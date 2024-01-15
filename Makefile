run:
	go run cmd/app/main.go
start:
	go run cmd/app/main.go
watch:
	air -d
build:
	go build -o api cmd/app/main.go
tidy:
	go mod download && go mod tidy
test:
	go test ./tests -v -cover
lint:
	golangci-lint run -v