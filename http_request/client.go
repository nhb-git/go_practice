package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
    url := flag.String("url", "", "request url")
    flag.Parse()
    count := 0

    for (count < 100) {
        go request_url(*url)
        count += 1
    }
}

func request_url(url string) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("err: ", err.Error())
    }
    fmt.Println("resp:", resp)
}