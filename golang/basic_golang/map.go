package main
import("fmt")

func main(){

	//di map keynya bisa string/bebas, lalu diikut nama mahasiswa yaitu string
	mahasiswa := make(map[string]string)//gunakan make untuk membuat slice

	//keynya berupa string, key = nim, lalu nama mhs
	mahasiswa["1234"] = "suharyadi"
	mahasiswa["4321"] = "erina"

	fmt.Println(mahasiswa)

	fmt.Println(mahasiswa["1234"])
	fmt.Println(mahasiswa["4321"])

	//nim = index, vMhs value menampung data, mahasiswa asal data arr nya
	for nim, vMhs := range mahasiswa {
		fmt.Println("nim mhs =", nim, "nama mahasiswa =", vMhs)
	}

	fmt.Println("----------------------")

	//map didalam map

	//tanpa menggunakan make
	mahasiswa2 := map[string]string{
		"111":"eko",
		"222":"eka",
	}
	fmt.Println(mahasiswa2)

	//nested map
	mahasiswa3 := map[string]map[string]string{
		"1234" : map[string]string{
			"name" : "sumanto",
			"tlp" : "098765",
		},
		"5678" : map[string]string{
			"name" : "sumanti",
			"tlp" : "234234",
		},
	}

	//delete data arr
	//diikuti nama array mapnya, lalu nama keynya
	//delete(mahasiswa3, "1234")

	fmt.Println(mahasiswa3)//panggil semua data mahasiswa 3
	fmt.Println(mahasiswa3["1234"])//panggil dengan nim 1234 saja
	fmt.Println(mahasiswa3["1234"]["name"])//panggil dengan nim 1234 dan nama mhs nya

	fmt.Println("----------------------")

	//perulangan array nested/map dalam map
	
	//perulangan pertama berdasarkan jumlah nim mahasiswa
	for ind, tmhs := range mahasiswa3 {
		fmt.Println("mahasiswa dgn nim = ", ind)//output nim mahasiswa
		//dikuti perulangan data nama dan tlp, 1 nim terdapat 2 data
		for inddua, tmhsdua := range tmhs {
			fmt.Println(inddua," = ", tmhsdua)//output name dan tlp
		}
	}

}