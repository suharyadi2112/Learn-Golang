package main 

import("fmt")

type mahasiswa struct{
	nim int `json:"nim"`
	name string `json:"name"`
	address string `json:"address"`
}

//embeded struct/mengambil struct mahasiswa kedalam struct matkul
type matkul struct{
	sks int `json:"sks"`
	matkul_name string `json:"matkul_name"`
	jenis string `json:"jenis"`
	mahasiswa
}

type dosen_pem struct{
	nip int `json"nip"`
	name string `json:"name"`
	jenis string `json:"jenis"`
	matkul
}

func main(){

	//Implementasi Struct Mahasiswa
	s := mahasiswa{
		nim : 3311401041,
		name : "Suharyadi",
		address : "Bumijo Kulon",
	}

	fmt.Println(s)
	fmt.Println("---------------------------------")
	//Implementasi Struct Matkul dengan Embeded dari Struct Mahasiswa 
	b := matkul{
		sks : 10,
		matkul_name : "daspro",
		jenis : "Pemrograman",
		mahasiswa : s,
	}

	fmt.Println(b)
	fmt.Println("---------------------------------")
	//Implementasi Struct dosem_pem dengan Embeded dari Struct matkul 
	c := dosen_pem{
		nip : 112312312,
		name : "eva liata",
		jenis : "penguji",
		matkul : b,
	}

	fmt.Println(c)
	fmt.Println("---------------------------------")


	//Implementasi Struct dengan properti sama antara 2 struck, 
	//Dalam kasus dibawah, struct dosen_pem memiliki jenis yang sama dengan matkul
	cek_dp := dosen_pem{}

	cek_dp.jenis = "penguji";//dari struct dosen_pem
	cek_dp.matkul.jenis = "Database";//dari struct matkul

	fmt.Println(cek_dp.jenis)
	fmt.Println(cek_dp.matkul.jenis)

}