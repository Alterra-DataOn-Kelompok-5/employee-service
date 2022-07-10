run:
	clear && go run main.go

test-all:
	clear && go test ./... -coverprofile=cover.out -p 1 && go tool cover -func=cover.out

show-test-result:
	go tool cover -html=cover.out

swagger:
	sudo docker run --rm -it --user $(shell id -u):$(shell id -g) -e GOPATH=$(HOME)/go:/go -e GOCACHE=/tmp -v $(HOME):$(HOME) -w $(shell pwd) quay.io/goswagger/swagger generate spec -o ./swagger.yml --scan-models