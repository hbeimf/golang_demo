package controller


import (
    "fmt"
    "net/http"
)



// curl -d "key1=val111&userId=1034285" "127.0.0.1:8080/post"
// post 请求
func (this *Controller) PostHandler(w http.ResponseWriter, req *http.Request) {
    defer req.Body.Close()

    post_val1 := req.PostFormValue("key1")
    fmt.Println(post_val1)

    fmt.Println("ok")
    w.Write([]byte("post_handler"))
}

