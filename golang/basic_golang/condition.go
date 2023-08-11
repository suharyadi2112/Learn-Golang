package main 

import(
	"fmt"
)

func main(){

	//for dan if else
	for a := 1; a <= 30; a++{

		if a%2 == 1{
			fmt.Println(a,"ganjil")
		}else{
			fmt.Println(a, "genap")
		}

	}

	//switch
	b := 2
	switch b {
	case 1:
		fmt.Println("satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	default:
		/* code */
		return
	}
	
	//break dan continue
	for c := 1; c <= 50; c++{
		//skip untuk angka 30
		if c == 30 {
			continue
		}
		//break atau stop diangka 40
		if c == 40{
			break
		}
		fmt.Println(c, "b and c")
	}
	
}