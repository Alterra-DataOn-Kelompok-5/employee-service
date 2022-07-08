run:
	clear && go run main.go

test-all:
	clear && go test ./... -coverprofile=cover.out -p 1 && go tool cover -func=cover.out

show-test-result:
	go tool cover -html=cover.out