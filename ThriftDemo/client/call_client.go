package main

import (
    "git.apache.org/thrift.git/lib/go/thrift"
    "net"
    "fmt"
    "golang_demo/ThriftDemo/call"
    "log"
)

const (
    HOST = "localhost"
    // PORT = "8088"
    // HOST = "127.0.0.1"
    PORT = "9009"
)

// https://blog.csdn.net/conggova/article/details/80696923

func main()  {
    tSocket, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
    if err != nil {
        log.Fatalln("tSocket error:", err)
    }

    transportFactory := thrift.NewTBufferedTransportFactory(1024*1024) 
    // 传输器
    // transportFactory := thrift.NewTBufferedTransportFactory(10000000)
    // transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
    transport := transportFactory.GetTransport(tSocket)
    // 传输协议
    protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

    client := call.NewCallServiceClientFactory(transport, protocolFactory)

    if err := transport.Open(); err != nil {
        log.Fatalln("Error opening:", HOST + ":" + PORT)
    }
    defer transport.Close()


    msg := call.Message{ID: 10000, Text:"hello,world call!"}
    d, err := client.Call(&msg)
    fmt.Println(d.Text)
    fmt.Println(d.ID)
    fmt.Println(d)

}
