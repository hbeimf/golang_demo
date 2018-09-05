# GOPATH := $(shell pwd)
MAIN := xdb

all: run

# server: server.go
# 	@GOPATH=$(GOPATH) go get -d
# 	@GOPATH=$(GOPATH) go build -o $@


xdb: $(MAIN).go
	go get -d
	go build -o $@

run:
	# ./GinDemo
	go get -d
	# go run ./GinDemo.go
	go run ./$(MAIN).go


cc:
	rm -rf ./$(MAIN)


# https://github.com/getlantern/lantern
start_proxy:
	/usr/bin/lantern &

set_proxy:
	git config --global http.proxy http://127.0.0.1:1080 
	git config --global https.proxy https://127.0.0.1:1080

unset_proxy: 
	git config --global --unset http.proxy 
	git config --global --unset https.proxy

show_proxy:
	git config http.proxy
