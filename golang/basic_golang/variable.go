package main

import "fmt"
//varible global agar bisa diakses kesemua function
var globe string = "Variable Globel"

func main(){

	var a string = "Hello World"
	fmt.Println(a)
	a ="Hello Go"
	fmt.Println(a)

	var ab string
	fmt.Println(ab)

	ab = "Isi ab"
	fmt.Println(ab)

	//tidak perlu menggunakan tipe data
	ac :=  "Persingkat Variable"
	fmt.Println(ac)

	ad := 10
	fmt.Println(ad)

	//varible constant, varible yang tidak bisa diubah nilainya
	const af string = "variable constant"
	fmt.Println(af)

}
//function ini bisa mengakses varible globe karena globe di set untuk bisa di akses kesemua function tes
func x(){
 	fmt.Println(globe)
}
