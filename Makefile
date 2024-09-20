# Makefile
export GO111MODULE=on
export DOCKER_BUILDKIT=1

APP_NAME=hello
DOCKER_COMPOSE=docker compose
GOTEST=go test

run:
	go run cmd/main.go

dev:
	air -c .air.toml

develop:
	$(DOCKER_COMPOSE) -f docker-compose.yml up -d --build --wait

stop-develop:
	$(DOCKER_COMPOSE) stop

down-develop:
	$(DOCKER_COMPOSE) down --volumes

logs:
	$(DOCKER_COMPOSE) logs -f

build-go:
	rm -f build/$(APP_NAME)
	go build -o build/$(APP_NAME) cmd/main.go

test:
	$(info Running all Go tests...)
	$(GOTEST) -v ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy

clean:
	docker image prune -f

.PHONY: run dev develop stop-develop down-develop logs build-go test fmt tidy clean