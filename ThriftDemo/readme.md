Thrift version 0.9.3

php, golang, erlang, rpc通信demo

golang 的坑真多，

参考：

https://studygolang.com/articles/9607
gopath : /golang
cd /golang
go get git.apache.org/thrift.git/lib/go/thrift

在GOPATH [/golang/src] 下建项目 

thrift -r --gen go example.thrift

这个要操作，不然会报错，版本不兼容
cd /golang/src/git.apache.org/thrift.git/lib/go/thrift
git checkout 0.9.3

