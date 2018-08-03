package main

import (
    "flag"
    "fmt"
)


// Go解析命令行传入参数
// https://blog.csdn.net/u011304970/article/details/71213359
// go run demo.go -src="123" -level=4 -memo="567"

func main() {
    var src string
    flag.StringVar(&src, "src", "", "source file")
    var level *int
    level = flag.Int("level", 0, "debug level")
    var memo string
    flag.StringVar(&memo, "memo", "", "the memory")
    flag.Parse()
    flag.Usage()
    fmt.Printf("src=%s, level=%d, memo=%s\n", src, *level, memo)
}