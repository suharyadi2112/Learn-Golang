package main

import (
    "crypto/rand"
    "fmt"
    "io"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func main() {
    cek := GenerateSetPin(6)

    fmt.Println(cek)
}
//generate set pin
func GenerateSetPin(Panjang int)(res string) {
    b := make([]byte, Panjang)
    n, err := io.ReadAtLeast(rand.Reader, b, Panjang)
    if n != Panjang {
        panic(err)
    }
    for i := 0; i < len(b); i++ {
        b[i] = table[int(b[i])%len(table)]
    }
    res = string(b)
    return res
}
