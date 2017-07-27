GOPATH := $(shell pwd)

all: HttpServer

test: Test

Test:
	@GOPATH=$(GOPATH) go get -d
	@GOPATH=$(GOPATH) go test -v ./model/Redis_test.go ./model/Redis.go
#	@GOPATH=$(GOPATH) go test -v ../model/ElasticSearch_test.go ../model/ElasticSearch.go
	@GOPATH=$(GOPATH) go test -v ./model/Mysql_test.go ./model/Mysql.go

HttpServer:
	@GOPATH=$(GOPATH) go get -d
	@GOPATH=$(GOPATH) go build -o $@

clean:
	@GOPATH=$(GOPATH) go clean
	@${RM} -r pkg/ src/

.PHONY: HttpServer
