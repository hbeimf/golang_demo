package controller


import (
    "fmt"
    "net/http"
)



// curl http://127.0.0.1:8080/get?key=val1
// curl http://127.0.0.1:8080/get?key=val123
// get请求
func (this *Controller) GetHandler(w http.ResponseWriter, req *http.Request) {
    defer req.Body.Close()

    query := req.URL.Query()

    val := query["key"][0]
    fmt.Println(val)


    // redis demo
    this.redis.RedisSet("key1122", "val1XXXXX")
    val1, err := this.redis.RedisGet("key1122")
    if err == nil {
        fmt.Println(val1)
    }

    // mysql demo
    Sql := "select * from users limit 10"
    rows := this.mysql.MysqlGet(Sql)
    for k, v := range rows {
        fmt.Printf("k=%v, v=%v\n", k, v)

        fmt.Printf("k=%v, id=%v\n", k, v["id"])
        fmt.Printf("k=%v, username=%v\n", k, v["username"])
        fmt.Printf("k=%v, email=%v\n", k, v["email"])
    }



    // fmt.Println("ok")
    // w.Write([]byte("get_handler"))
    w.Write([]byte(val))

}
