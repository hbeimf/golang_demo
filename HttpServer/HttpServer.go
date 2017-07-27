package main

// doc
// http://blog.csdn.net/su_sai/article/details/51176753
// http 状态码
// http://blog.csdn.net/liangguangchuan/article/details/54575001
//
// http://studygolang.com/articles/2680

// go version go1.8.3 linux/amd64

// redis
// elasticsearch
// ssdb
// rabbitmq


import (
    "fmt"
    "flag"
    "net/http"
    "./controller"
)

// 设置全局配置变量，并带默认值
var (
    globalHttpHost = flag.String("globalHttpHost", "127.0.0.1:8080", "http服务监听主机端口")
)

func main() {
    flag.Parse()

    ctrl := controller.GetControllerInstance()

    // 注册控制器函数
    http.HandleFunc("/get", ctrl.GetHandler)
    http.HandleFunc("/post", ctrl.PostHandler)

    // websocket echo demo
    http.HandleFunc("/websocket", ctrl.WebsocketHandler)

    err := http.ListenAndServe(*globalHttpHost, nil)
    if err != nil {
        fmt.Println("err")
    }
}






