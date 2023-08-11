package main

import (
	"fmt"
	"math"
)

//interface tipe data yang bernilai nil, baru bisa digunakan jika sudah memiliki isi
type hitung interface {
	luas() float64     //function luas dijabarkan diinterface
	keliling() float64 //function keliling dijabarkan diinterface
}

//----ini bagian struct
type lingkaran struct {
	diameter float64
}

type persegi struct {
	sisi float64
}

//lingkaran
func (l lingkaran) jariJari() float64 { // l merupakan alias dari lingkaran, yang lalu akses untuk diameter didalamnya
	return l.diameter / 2 //l diameter dari l lingkaran
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

//persegi
func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}
func main() {
	var bangunDatar hitung //penjabaran hitung dari interface sebelumnya

	bangunDatar = persegi{10.0} //masukan value kestruct persegi{value}
	fmt.Println("===== persegi")
	fmt.Println("luas      :", bangunDatar.luas())     //var bangunDatar langsung akses method luas()
	fmt.Println("keliling  :", bangunDatar.keliling()) //var bangunDatar langsung akses method keliling()

	bangunDatar = lingkaran{14.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())
	fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())
}
