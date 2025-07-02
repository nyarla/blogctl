all: build

build:
	go build -o blogctl ./...

tidy:
	go mod tidy

test:
	go test -v ./...

lint:
	golangci-lint run

podman:
	env DOCKER_HOST=unix:///run/user/$$(id -u)/podman/podman.sock $(CMD)

act-pull_request:
	@$(MAKE) podman CMD="act pull_request"
