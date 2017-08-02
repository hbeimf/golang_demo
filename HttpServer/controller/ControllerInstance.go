package controller


import (
    // "flag"
    "../model"
)


type Controller struct {
    redis *model.Redis
    mysql *model.Mysql
}

var ControllerInstance *Controller

func init() {
    redis := model.RedisClient
    mysql := model.MysqlClient
    ControllerInstance = &Controller{
        redis,
        mysql,
    }
}


func GetControllerInstance() *Controller {
    return ControllerInstance
}

