# GOPATH := $(shell pwd)

all: GinDemo

# server: server.go
# 	@GOPATH=$(GOPATH) go get -d
# 	@GOPATH=$(GOPATH) go build -o $@


GinDemo: GinDemo.go
	go get -d
	go build -o $@

run:
	# ./GinDemo
	go get -d
	go run ./GinDemo.go

cc:
	rm -rf ./GinDemo
