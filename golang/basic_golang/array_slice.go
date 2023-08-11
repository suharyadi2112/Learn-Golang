package main

import(
	"fmt"
)

func main(){
	//perbedaan array biasa dengan slice yaitu, slice nomor/indexnya tidak ada
	buah := []string{//[]tidak ada nomor

		"apel",
		"jeruk",
		"durian",
		"markisa",

	}
	fmt.Println(buah)

	fmt.Println("------")

	buah2 := make([]string, 4)//membuat slice dengan make()

	//data ini berupa array yang ditampung dalam slice buah2
	buah2[0] = "ubi" 
	buah2[1] = "melon"  
	buah2[2] = "rambutan"
	buah2[3] = "mangga"

	fmt.Println(buah2)

	buah2[2] = "pisang"//merubah isi slice buah2 index 2 menjadi pisang

	fmt.Println(buah2)

	fmt.Println("------")

	//menambah isi slice menggunakan fungsi append

	slice2 := buah2//slice2 menampung isi dari slice buah2
	slice2 = append(slice2, "nanas","pepaya")//slice ditambah menggunakan  append

	fmt.Println(slice2)//slice2 sudah bertamaah



	fmt.Println("------")

	slice3 := make([]string, 10)//slice 3 untuk mengcopy data dari array buah2
	//jika berlebih contoh 10, maka sisa nya kosong

	copy(slice3, buah2)//copy digunakan untuk duplicate/copy data dari array buah 2

	fmt.Println(slice3)

}