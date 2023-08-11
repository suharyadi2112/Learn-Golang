package main
import (
  "fmt"
  "runtime"
  "errors"
)


var a bool = false

// Goroutine mirip dengan thread, tapi sebenarnya bukan. Sebuah native thread bisa berisikan sangat banyak goroutine. Mungkin lebih pas kalau goroutine disebut sebagai mini thread,  Eksekusi goroutine bersifat asynchronous, menjadikannya tidak saling tunggu dengan goroutine lain.
func main() {
  // Fungsi ini digunakan untuk menentukan jumlah core atau processor yang digunakan dalam eksekusi program.
  runtime.GOMAXPROCS(2)

  for i := 0; i < 5; i++ {
    print(i)
  }

}

//Penerapan Goroutine
func print (i int)  {

  fmt.Println(a)
  if !a {
      a = false
      
      err = errors.New("math: square root of negative number")
      
      if err != nil {
        fmt.Println(err.Error())
      }else{
        a = true
      }
    
  }

}

