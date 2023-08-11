package main

import (
	"fmt"
)

func main() {
	//array
	var name [3]string

	name[0] = "satu"
	name[1] = "dua"
	name[2] = "tiga"

	fmt.Println(name)    //tampil semua arr
	fmt.Println(name[1]) //panggil 1 saja

	fmt.Println("------")

	for i := 0; i < len(name); i++ { //menggunakan perulangan
		fmt.Println(name[i])
	}
	fmt.Println(len(name)) //len untuk hitung total arr

	fmt.Println("------")

	//menggunakan foreach
	//index untuk indek foreach (penamaan bebas)
	//value untuk menampung element dari name (penamaan bebas)
	//name sendiri sumber array (sesuai nama array)
	for index, value := range name {
		fmt.Println("indexnya", index, "valuenya", value)
	}

	fmt.Println("------")

	for _, value := range name { //index diganti underscore _ jika tidak digunakan
		fmt.Println("valuenya", value)
	}

	//Slice mengambil referensi dari array yang dideklerasikan
	fmt.Println("------")

	buah := [4]string{
		"apel", "jeruk", "jambu", "pear",
	}

	slice1 := buah[2:]  //mengambil potongan dari index 2 sampai akhir
	slice2 := buah[:2]  //mengambil potongan dari awal index 0 sampai index 2
	slice3 := buah[0:1] //mengambil potongan index 0 sampai index 2

	buah[1] = "durian" //merubah isi array index 1

	fmt.Println(slice1, slice2, slice3, buah)

	fmt.Println("------")

}
