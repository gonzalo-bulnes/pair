define GOLANG_CMD
	docker run --rm -it \
		--volume "$$GOPATH":/gopath \
		--volume "$$(pwd)":/app \
		--env "GOPATH=/gopath" \
		--workdir /app \
		golang:1.12-alpine
endef

.PHONY: all
all: clean build

.PHONY: build
build: pair-linux-amd64 pair-darwin-amd64

.PHONY: clean
clean:
	@rm -f pair

.PHONY: install-for-testing
install-for-testing:
	go get -t ./...
	go get github.com/redbubble/go-passe

pair-linux-amd64: *.go **/*.go
	@${GOLANG_CMD} \
		sh -c 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a --installsuffix cgo --ldflags="-s" -o pair-linux-amd64 ./cmd'

pair-darwin-amd64: *.go **/*.go
	@${GOLANG_CMD} \
		sh -c 'CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a --installsuffix cgo --ldflags="-s" -o pair-darwin-amd64 ./cmd'

.PHONY: test
test:
	go test -json ./... | go-passe
