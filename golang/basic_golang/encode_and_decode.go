package main

import (
  "encoding/base64"
  "fmt"
)

func main() {
    var data = "john wick"

    //bentukan jika string data di casting ke []byte
    fmt.Println("------Penerapan Fungsi EncodeToString() & DecodeString()------")
    fmt.Println([]byte(data))
    fmt.Println("------------")

    //encoded string data wajib di casting ke []byte terlebih dahulu
    fmt.Println("------------")
    var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println("encoded:", encodedString)
    fmt.Println("------------")

    //decodestring hari dari encode string, jika string biasa maka error
    var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
    var decodedString = string(decodedByte)
    fmt.Println("decoded:", decodedString)
    fmt.Println("------------")


    fmt.Println("-----Penerapan Fungsi Encode() & Decode()-------")
    //Fungsi base64.StdEncoding.EncodedLen(len(data)) menghasilkan informasi lebar variable data ketika sudah di-encode. Nilai tersebut kemudian ditentukan sebagai lebar alokasi tipe []byte pada variabel encoded yang nantinya digunakan untuk menampung hasil encoding.
    var encodedDua = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
    base64.StdEncoding.Encode(encodedDua, []byte(data))

    //hasil encoded
    var encodedStringDua = string(encodedDua)
    fmt.Println(encodedStringDua)

    //Fungsi base64.StdEncoding.DecodedLen() memiliki kegunaan sama dengan EncodedLen(), hanya saja digunakan untuk keperluan decoding.
    var decodedDua = make([]byte, base64.StdEncoding.DecodedLen(len(encodedDua)))
    var _, err = base64.StdEncoding.Decode(decodedDua, encodedDua)

      if err != nil {
        fmt.Println(err.Error())
    }
    //hasil Decode
    var decodedStringDua = string(decodedDua)
    fmt.Println(decodedStringDua)

    fmt.Println("-----Encode & Decode Data URL-------")
    // /Khusus encode data string yang isinya merupakan URL, lebih efektif menggunakan URLEncoding dibandingkan StdEncoding.
    var dataDua = "https://kalipare.com/"

    var encodedStringTiga = base64.URLEncoding.EncodeToString([]byte(dataDua))
    fmt.Println(encodedStringTiga)

    var decodedByteTiga, _ = base64.URLEncoding.DecodeString(encodedStringTiga)
    var decodedStringTiga = string(decodedByteTiga)
    fmt.Println(decodedStringTiga)

}
