package main
import (
  "fmt"
  "runtime"
)
//Penerapan Goroutine
func print (n int, pesan string)  {
  for i := 0;i < n; i++ {
    fmt.Println((i + 1), pesan)
  }
}
// Goroutine mirip dengan thread, tapi sebenarnya bukan. Sebuah native thread bisa berisikan sangat banyak goroutine. Mungkin lebih pas kalau goroutine disebut sebagai mini thread,  Eksekusi goroutine bersifat asynchronous, menjadikannya tidak saling tunggu dengan goroutine lain.
func main() {
  // Fungsi ini digunakan untuk menentukan jumlah core atau processor yang digunakan dalam eksekusi program.
  runtime.GOMAXPROCS(2)

  go print(100, "halo")
  print(100, "apa kabar")

  var input string
  // Fungsi ini akan meng-capture semua karakter sebelum user menekan tombol enter, lalu menyimpannya pada variabel.
  //fungsi fmt.Scanln(). Fungsi tersebut bisa menampung parameter bertipe interface{} berjumlah tak terbatas. 
  fmt.Scanln(&input)
}
