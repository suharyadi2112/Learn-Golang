package main

import (
    // "encoding/json"
    "fmt"
    // "github.com/jmcvetta/napping"
    // "log"
    "net/http"
    "io/ioutil"
    "bytes"
)

func main() {
    url := "http://localhost:9999/newuser"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`
        { 
            "name": "Jhonson", 
            "email": "Jhonson@gmail.com"
        }`)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

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
