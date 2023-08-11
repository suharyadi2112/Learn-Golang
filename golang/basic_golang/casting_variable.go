package main

import (
  "fmt"
  "strings"
  "reflect"
)

type person struct {
  name string
  age int
}

func main() {
  //Casting variable interface kosong
  var cekSatu interface{}
  cekSatu = 10

  //varibale cekSatu perlu dicasting dahulu ke int, dikarenakan nilai cekSatu merupakan string
  cekDua := cekSatu.(int)*10
  fmt.Println(cekDua)

  //Join menggunakan string.Join
  cekSatu = []string{"apple", "manggo", "banana"}//Teknik casting pada interface disebut dengan type assertions.
  var gruits = strings.Join(cekSatu.([]string), ", ")
  fmt.Println(gruits, "is my favorite fruits")

//--------------------------//
  //Casting Variabel Interface Kosong Ke Objek Pointer
  var secretDua interface{} = &person{name: "wick", age: 27}
  // Cara casting dari interface{} ke struct pointer adalah dengan menuliskan nama struct-nya dan ditambahkan tanda asterisk (*) di awal, contohnya seperti secret.(*person). Setelah itu barulah nilai asli bisa diakses.
  name := secretDua.(*person).name
  fmt.Println(name)

//---------------------//
  //Kombinasi Slice, map, dan interface{}
  var dataLima = []map[string]interface{}{
    {"nama" : "Suharyadi", "almt" : "Bumjo"},
    {"nama" : "Owl", "almt" : "Atas Pohon"},
  }
  for _, ck := range dataLima {
    fmt.Println(ck["nama"])
  }

  cekRef := reflect.ValueOf(dataLima)//cek tipe dan value data menggunakan reflect
  fmt.Println(cekRef.Type())//cek tipe dan value data menggunakan reflect
}
