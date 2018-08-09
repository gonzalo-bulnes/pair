.PHONY: all
all: clean build

.PHONY: build
build: pair

.PHONY: clean
clean:
	@rm -f pair

.PHONY: install-for-testing
install-for-testing:
	go get -t ./...
	go get github.com/redbubble/go-passe

pair:
	@docker run --rm -it \
		--volume "$$GOPATH":/gopath \
		--volume "$$(pwd)":/app \
		--env "GOPATH=/gopath" \
		--workdir /app \
		golang:1.9-alpine \
		sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o pair'

.PHONY: test
test:
	go test -json ./... | go-passe
