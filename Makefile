all: build

build:
	go build -o blogctl ./...

tidy:
	go mod tidy

test:
	go test -v ./...

podman-test:
	env DOCKER_HOST=unix:///run/user/$$(id -u)/podman/podman.sock act pull_request
