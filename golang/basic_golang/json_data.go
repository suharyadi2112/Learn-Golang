package main

import (
	"encoding/json"
	"fmt"
)

//User struct digunakan untuk menampung hasil decode json string
//Proses decode sendiri dilakukan lewat fungsi json.Unmarshal()
type User struct {
  Name string `json:"name"`// Tag `json:"name"` digunakan untuk mapping informasi json ke property yang bersangkutan.
  Age int
}

func main() {
//---------- Decode JSON Ke Variabel Objek Struct------------//
  fmt.Println("------------Decode JSON Ke Variabel Objek Struct----------------")
  jsonString := `{"Name": "john wick", "Age": 27}`
  jsonData := []byte(jsonString)//casting jsonString ke []byte

  //declare variable dari struct
  var data User
  // Fungsi unmarshal hanya menerima data json dalam bentuk []byte
  // argument ke-2 fungsi unmarshal harus diisi dengan pointer dari objek yang nantinya akan menampung hasilnya.
  var err = json.Unmarshal(jsonData, &data)

    if err != nil {//check error dari json.Unmarshal
        fmt.Println(err.Error()) 
        return
    }

  fmt.Println("Cek Nama : ", data.Name)
  fmt.Println("Cek Umur : ", data.Age)
//------------ Decode JSON Ke Variabel Objek Struct------------//


//------Decode JSON Ke map[string]interface{} & interface{}-----//
  fmt.Println("------------Decode JSON Ke map[string]interface{} & interface{}----------------")
  var data1 map[string]interface{}
  json.Unmarshal(jsonData, &data1)

  fmt.Println("user :", data1["Name"])
  fmt.Println("age  :", data1["Age"])


  //Variabel bertipe interface{} juga bisa digunakan untuk menampung hasil decode. Dengan catatan pada pengaksesan nilai property, harus dilakukan casting terlebih dahulu ke map[string]interface{}.
  var data2 interface{}
  json.Unmarshal(jsonData, &data2)
  var decodedData = data2.(map[string]interface{})
  fmt.Println(decodedData)
  fmt.Println(decodedData["Name"])
  fmt.Println(decodedData["Age"])
//------Decode JSON Ke map[string]interface{} & interface{}-----//

//------Decode Array JSON Ke Array Objek-------//
  fmt.Println("----------Decode Array JSON Ke Array Objek----------------")
  var jsonStringDua = `[
      {"Name": "john wick", "Age": 27},
      {"Name": "ethan hunt", "Age": 32}
  ]`
  ngebyteJson := []byte(jsonStringDua)
  var Cekuser []User
  // json.Unmarshal([]byte(jsonStringDua), &data2)//cara 1 langsung ngebyte dalam unmarshal
  json.Unmarshal(ngebyteJson, &Cekuser)//cara 2 // parameter ke 2 untuk tampung data

  fmt.Println(Cekuser[0].Name)
  fmt.Println(Cekuser[1].Name)
//------Decode Array JSON Ke Array Objek-------//


//------Encode Objek Ke JSON String-------//
  fmt.Println("----------Encode Objek Ke JSON String----------------")
  // Fungsi json.Marshal digunakan untuk encoding data ke json string. Sumber data bisa berupa variabel objek cetakan struct, map[string]interface{}, atau slice.

  var object = []User{{"john wick", 27}, {"ethan hunt", 32}}//data slice
  fmt.Println(object)//keluaran data slice

  var jsonDataTiga, errt = json.Marshal(object)// convert to Marshal []byte

  if errt != nil {
      fmt.Println(errt.Error())
      return
  }
  // Pada contoh berikut, data slice struct dikonversi ke dalam bentuk json string. Hasil konversi berupa []byte, casting terlebih dahulu ke tipe string agar bisa ditampilkan bentuk json string-nya.
  var jsonStringTiga = string(jsonDataTiga)//[]byte marshal to string

  fmt.Println(jsonStringTiga)


  fmt.Println("----------Get Property Dari Encode----------------")
  //get property dari encode sebeleumnya
  TigaByte := []byte(jsonDataTiga)
  var CekUserTiga []User

  json.Unmarshal(TigaByte, &CekUserTiga)//parameter ke2 untuk tampung data, kasus ini ke struct

  //looping berdasarkan jumlah data
  for g := 0; g < len(CekUserTiga); g++ {
    fmt.Println("Name : ", CekUserTiga[g].Name, ", Age : ", CekUserTiga[g].Age)//get property
  }
  // fmt.Println(CekUserTiga[c].Name)


//------Encode Objek Ke JSON String-------//

}
