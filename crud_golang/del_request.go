package main

import (
    // "encoding/json"
    "fmt"
    // "github.com/jmcvetta/napping"
    // "log"
    "net/http"
    "io/ioutil"
)

func main() {
    //url untuk delete , harus menggunakan parameter
    url := "http://localhost:9999/deluser/3"
    fmt.Println("URL:>", url)//url yang dituju

    req, err := http.NewRequest("DELETE", url, nil)//request HTTP
    req.Header.Set("X-Custom-Header", "myvalue")//header
    req.Header.Set("Content-Type", "application/json")//header

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
