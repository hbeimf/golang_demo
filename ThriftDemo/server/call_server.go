package main

import (
    "ThriftDemo/call"
    "strings"
    "git.apache.org/thrift.git/lib/go/thrift"
    "fmt"
    "log"
)

type CallImpl struct {}

func (fdi *CallImpl) Call(msg *call.Message) (r *call.Message, err error){
    var rData call.Message
    rData.Text = strings.ToUpper(msg.Text)
    rData.ID = msg.ID

    fmt.Println(msg.Text)
    fmt.Println(msg.ID)

    return &rData, nil
}

const (
    HOST = "localhost"
    // PORT = "8088"
    PORT = "9009"
)

func main() {

    handler := &CallImpl{}
    processor := call.NewCallServiceProcessor(handler)
    serverTransport, err := thrift.NewTServerSocket(HOST + ":" + PORT)
    if err != nil {
        log.Fatalln("Error:", err)
    }

    transportFactory := thrift.NewTBufferedTransportFactory(1024*1024) 
    // 传输器 https://blog.csdn.net/conggova/article/details/80696923
    // transportFactory := thrift.NewTBufferedTransportFactory(10000000)
    // transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
    protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

    server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
    fmt.Println("Running at:", HOST + ":" + PORT)
    server.Serve()
}
