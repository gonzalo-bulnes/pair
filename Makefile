.PHONY: all
all: clean build

.PHONY: build
build: pair

.PHONY: clean
clean:
	@rm -f pair

pair:
	@docker run --rm -it \
		--volume "$$GOPATH":/gopath \
		--volume "$$(pwd)":/app \
		--env "GOPATH=/gopath" \
		--workdir /app \
		golang:1.9-alpine \
		sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o pair'
