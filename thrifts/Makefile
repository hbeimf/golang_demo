GOPATH := $(shell pwd)

all: client


# server: server.go
# 	@GOPATH=$(GOPATH) go get -d
# 	@GOPATH=$(GOPATH) go build -o $@


client: client.go
	@GOPATH=$(GOPATH) go get -d
	@GOPATH=$(GOPATH) go build -o $@

run:
	./client

cc:
	rm -rf ./client

clean:
	@GOPATH=$(GOPATH) go clean
	@${RM} -r pkg/ src/

.PHONY: HttpServer