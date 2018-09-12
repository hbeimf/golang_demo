运行demo

```
$ cd HttpServer

$ make

$ make test


$ ./HttpServer


$ go run WebsocketClient.go
```

依赖环境 ：

```
rabbitmq
redis
mysql
```



mysqldump -uroot -p test users > /web/golang_demo/test_users.sql


```
$ mkdir -p /golang/src/golang_demo

ERL_HOME=/usr/local/erlang
PATH=$ERL_HOME/bin:$PATH
export ERL_HOME PATH

PATH=$PATH:/usr/local/php/bin
export PATH

PATH=/usr/local/mysql/bin:$PATH
export PATH

export GOPATH=/golang
export GOBIN=/usr/local/go/bin
export PATH=$PATH:$GOBIN



export JAVA_HOME=/usr/local/java/jdk1.8.0_181 
export JRE_HOME=${JAVA_HOME}/jre  
export CLASSPATH=.:${JAVA_HOME}/lib:${JRE_HOME}/lib  
export PATH=${JAVA_HOME}/bin:$PATH


go version go1.9.2 linux/amd64

```

