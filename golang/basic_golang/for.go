package main

import(

	"fmt"

)
func main(){

//perulangan increment dengan for dengan penjabaran, lebih dipisah
	a := 1
	for a < 10 {
		fmt.Println("Hello World", a)
	a++
	}
	fmt.Println("Selesai")
	
// -----------

// Perulangan decrement dengan for lebih clear
	for b := 10; b >= 1; b--{
			fmt.Println("Hello World 2", b)
	}
	fmt.Println("Selesai 2")

}