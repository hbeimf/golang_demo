package controller


import (
    "fmt"
    "flag"
    "net/http"
    "../model"
    "github.com/gorilla/websocket"
    "log"
)


var (
    globalRedisHost = flag.String("globalRedisHost", "127.0.0.1:6379", "redis服务监听主机端口")
)


type Controller struct {
    redis *model.Redis
}

var upgrader = websocket.Upgrader{} // use default options

var ControllerInstance *Controller

func init() {

    redis := model.NewRedisPool(*globalRedisHost, 0)
    ControllerInstance = &Controller{redis}
}


func GetControllerInstance() *Controller {
    return ControllerInstance
}



func (this *Controller) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
    c, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("upgrade:", err)
        return
    }
    defer c.Close()
    for {
        mt, message, err := c.ReadMessage()
        if err != nil {
            log.Println("read:", err)
            break
        }
        log.Printf("recv: %s", message)
        err = c.WriteMessage(mt, message)
        if err != nil {
            log.Println("write:", err)
            break
        }
    }
}



// curl http://127.0.0.1:8080/get?key=val1
// curl http://127.0.0.1:8080/get?key=val123
// get请求
func (this *Controller) GetHandler(w http.ResponseWriter, req *http.Request) {
    defer req.Body.Close()

    query := req.URL.Query()

    val := query["key"][0]
    fmt.Println(val)


    this.redis.RedisSet("key1122", "val1XXXXX")
    val1, err := this.redis.RedisGet("key1122")
    if err == nil {
        fmt.Println(val1)
    }

    // fmt.Println(val1)



    fmt.Println("ok")
    // w.Write([]byte("get_handler"))
    w.Write([]byte(val))

}

// curl -d "key1=val111&userId=1034285" "127.0.0.1:8080/post"
// post 请求
func (this *Controller) PostHandler(w http.ResponseWriter, req *http.Request) {
    defer req.Body.Close()

    post_val1 := req.PostFormValue("key1")
    fmt.Println(post_val1)

    fmt.Println("ok")
    w.Write([]byte("post_handler"))
}

