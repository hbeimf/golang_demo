GOPATH := $(shell pwd)

all: main

main: main.go 
	@GOPATH=$(GOPATH) go get -d
	@GOPATH=$(GOPATH) go build -o $@

run: main
	./main

cc:
	rm -rf ./main

clean:
	@GOPATH=$(GOPATH) go clean
	@${RM} -r pkg/ src/

.PHONY: HttpServer