package main
import (
  "fmt"
  "strings"
)
func main() {
  //----Fungsi strings.Contains()----//
  //Dipakai untuk deteksi apakah string (parameter kedua) merupakan bagian dari string lain (parameter pertama). Nilai kembaliannya berupa bool.
  var truee = strings.Contains("john wick", "wick")//hasil true
  var falsee = strings.Contains("john wick", "widck")//hasil false
  fmt.Println(truee)
  fmt.Println(falsee)

  //-----Fungsi strings.HasPrefix()------//
  //Digunakan untuk deteksi apakah sebuah string (parameter pertama) diawali string tertentu (parameter kedua).
  var isPrefix1 = strings.HasPrefix("john wick", "jo")
  fmt.Println(isPrefix1) // true
  var isPrefix2 = strings.HasPrefix("cek wick", "ce")
  fmt.Println(isPrefix2) // true
  var isPrefix3 = strings.HasPrefix("lewick", "ce")
  fmt.Println(isPrefix3) // false

  //Info lebih mengenai strings
  //https://dasarpemrogramangolang.novalagung.com/A-strings.html 
}
