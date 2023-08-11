package main 

import(
	"fmt"
)
//Pointer adalah reference atau alamat memori (hexadesimal)
func main(){

	//tidak bisa menampung nilai yang bukan pointer (reference memory)
	var name string = "Suharyadi"//ditandai dengan simbol arterisk
	var nameb *string = &name

	fmt.Println(*nameb)//Output Suharyadi
	fmt.Println(nameb)//OutPut 0xc00003e250


	//parameter pointer 

	var number int = 4
	//sebelum change menggunakan parammeter pointer
	fmt.Println(number)

	change(&number, 6)
	//sesudah change menggunakan parameter pointer
	fmt.Println(number)
}
func change(org *int, value int){
	*org = value //perubahan referensi nilai org(4), value(6)
}