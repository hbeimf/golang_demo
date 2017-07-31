package model

import (
    "testing"
)

func TestGet(t *testing.T) {
    val, err := RedisClient.RedisGet("key1122")
    t.Log("getXXX", val, err)
}

func TestSetGet(t *testing.T) {
    RedisClient.RedisSet("key100", "value_100")

    val, err := RedisClient.RedisGet("key100")
    t.Log("设置", val, err)
}


func TestLPush(t *testing.T) {
    listName := "list100"
    RedisClient.RedisLPush(listName, "value222")
    res, err := RedisClient.RedisRPop(listName)
    t.Log("队列左进右出", res, err)
}



func TestSetNX(t *testing.T) {
    RedisClient.RedisSetNX("key1000", "value_1000")

    RedisClient.RedisEXPIRE("key1000", 30)
}


func TestPubSub(t *testing.T) {
    channle := "test_channel"
    go RedisClient.RedisSub(channle)
    go RedisClient.RedisSub(channle)
    go RedisClient.RedisSub(channle)

    msg := "test msg !!!XXXXX"
    RedisClient.RedisPublish(channle, msg)

}



