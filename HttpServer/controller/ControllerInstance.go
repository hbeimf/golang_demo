package controller


import (
    // "flag"
    "../model"
)


type Controller struct {
    redis *model.Redis
}

var ControllerInstance *Controller

func init() {
    redis := model.RedisClient
    ControllerInstance = &Controller{redis}
}


func GetControllerInstance() *Controller {
    return ControllerInstance
}

