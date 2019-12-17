VERSION := $(shell git describe)
VERSION_WIDE := $(shell git describe)+$(shell date +'%_Y%m%_d')-$(shell git rev-parse --short=7 HEAD)
GO_PATH := $(shell go env GOPATH)

.PHONY: lint
lint:
	curl -sfL "https://install.goreleaser.com/github.com/golangci/golangci-lint.sh" | sh -s -- -b $(GO_PATH)/bin latest
	golangci-lint run

.PHONY: build
build:
	GOOS=linux go build -o watchman-linux-amd64-$(VERSION)/watchman -x -ldflags "-s -w -X main.Version=$(VERSION_WIDE)" cmd/watchman/main.go
	tar -czvf watchman-linux-amd64-$(VERSION).tar.gz watchman-linux-amd64-$(VERSION)
	GOOS=darwin go build -o watchman-darwin-amd64-$(VERSION)/watchman -x -ldflags "-s -w -X main.Version=$(VERSION_WIDE)" cmd/watchman/main.go
	tar -czvf watchman-darwin-amd64-$(VERSION).tar.gz watchman-darwin-amd64-$(VERSION)
	GOOS=windows go build -o watchman-windows-amd64-$(VERSION)/watchman -x -ldflags "-s -w -X main.Version=$(VERSION_WIDE)" cmd/watchman/main.go
	tar -czvf watchman-windows-amd64-$(VERSION).tar.gz watchman-windows-amd64-$(VERSION)


