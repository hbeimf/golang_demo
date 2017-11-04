GOPATH := $(shell pwd)

all: main

main: main.go ./stack/get_min.go
	@GOPATH=$(GOPATH) go get -d
	@GOPATH=$(GOPATH) go build -o $@

run:
	./main

cc:
	rm -rf ./main

clean:
	@GOPATH=$(GOPATH) go clean
	@${RM} -r pkg/ src/

.PHONY: HttpServer