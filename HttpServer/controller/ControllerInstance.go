package controller


import (
    "flag"
    "../model"
)


var (
    globalRedisHost = flag.String("globalRedisHost", "127.0.0.1:6379", "redis服务监听主机端口")
)


type Controller struct {
    redis *model.Redis
}

var ControllerInstance *Controller

func init() {
    flag.Parse()

    redis := model.NewRedisPool(*globalRedisHost, 0)
    ControllerInstance = &Controller{redis}
}


func GetControllerInstance() *Controller {
    return ControllerInstance
}

