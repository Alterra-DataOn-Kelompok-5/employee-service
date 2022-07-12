IMAGE_TAG_NAME=azkafr92/meeting-room-management-employee-service:latest
build:
	go build -o ./employee-service

down:
	docker compose down --rmi local --remove-orphans -v 

run:
	clear && go run main.go

show-test-result:
	go tool cover -html=cover.ou

stop:
	docker compose stop

test-all:
	clear && go test ./... -coverprofile=cover.out -p 1 && go tool cover -func=cover.out

up:
	docker compose up -d && docker compose start

build-image:
	docker build . -t ${IMAGE_TAG_NAME} -f docker/go/Dockerfile

push-image:
	docker push ${IMAGE_TAG_NAME}