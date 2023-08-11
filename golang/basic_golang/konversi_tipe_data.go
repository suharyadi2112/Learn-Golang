package main

import (
	"fmt"
	"strconv"
)

func main(){

	//mengubah tipe data int 32 ke 64
	var x int32 = 10
	var y int64 = int64(x)

	fmt.Println(y)

	//mengubah nilai berkoma float, menjadi int satuan
	var z float64 = 3.8
	var nm int64 = int64(z)

	fmt.Println(nm)

	//mengubah string menjadi int, contoh angka string "100", "200"
	var ac string = "100"

	ad, err := strconv.Atoi(ac)
	fmt.Println(ad, err)

	//mengubah Int menjadi string, contoh angka string menjadi angka int
	ag := strconv.Itoa(200)
	fmt.Println(ag)


}