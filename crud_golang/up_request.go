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
    url := "http://localhost:9999/UpArticle/3"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`
        { 
            "Title": "Hello 3 Update", 
            "desc": "Deksripsi Tiga Update", 
            "content": "Article Content Tiga Update"
        }`)

    req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
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
