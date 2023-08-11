package main

import (
	"fmt"
	"math"
)

//di interface hitung2d menampung 2 method
type hitung2d interface {
	luas() float64
	keliling() float64
}

//di interface hitung3d menampung 1 method
type hitung3d interface {
	volume() float64
}

//di interface hitung_again digunakan untuk EMBEDDED 2 interface menjadi 1
type hitungagain interface {
	hitung2d //interface 1
	hitung3d //interface 2
}
//hanyalah struct biasa
type kubus struct {
	sisi float64
}
//method hitung volume
func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

//method hitung luas
func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

//method hitung keliling
func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func main() {
	var bangunRuang hitungagain = &kubus{4}

	fmt.Println("===== kubus")
	fmt.Println("luas      :", bangunRuang.luas())
	fmt.Println("keliling  :", bangunRuang.keliling())
	fmt.Println("volume    :", bangunRuang.volume())
}
