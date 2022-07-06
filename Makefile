run:
	go run main.go

test-all:
	go test ./... -coverprofile=cover.out -v -p 1 && go tool cover -func=cover.out