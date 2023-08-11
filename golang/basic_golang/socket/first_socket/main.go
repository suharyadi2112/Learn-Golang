package main

import (
    "fmt"
    "github.com/googollee/go-socket.io"
    "net/http"
    "log"
)

func main() {
    server := socketio.NewServer(nil)
 
  
    server.OnConnect("/", func(s socketio.Conn) error {
        s.SetContext("")
        fmt.Println("connected:", s.ID())

        s.Join("chat")

         // send a welcome message to the client
        s.Emit("welcome", "Hello, client!")

        return nil
    })

    go server.Serve()
    defer server.Close()

    http.Handle("/socket.io/", server)
    http.Handle("/", http.FileServer(http.Dir("./asset")))
    log.Println("Serving at localhost:8000...")
    log.Fatal(http.ListenAndServe(":8000", nil))
}