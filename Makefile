lint: ###  run linters
	go vet ./...

test: ###  run tests
	go test ./...

client: ###  run client only
	cd cmd/client/ && go run main.go; cd ../../

server: ###  run server only
	cd cmd/server/ && go run main.go; cd ../../

compose-run: ###  run client and server in docker-compose
	docker-compose up

compose-rebuild: ###  rebuild images and run docker-compose
	docker-compose up --build

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
