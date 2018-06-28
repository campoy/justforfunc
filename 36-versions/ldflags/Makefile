dev:
	go build

HEAD = $(shell git rev-parse HEAD)
prod:
	go build -ldflags "-X main.version=$(HEAD)"