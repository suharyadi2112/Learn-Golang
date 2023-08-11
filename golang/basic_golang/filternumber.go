package main
import(
  "fmt"
  "strings"
)

func main(){
  cf := "87781112d238809"
  isNotDigit := func(c rune) bool {
    return c < '0' || c > '9'
  }
  b := strings.IndexFunc(cf, isNotDigit) == -1
  fmt.Println(b)
  //
  // var s []string
  // if b == false {
  //   for _, c := range cf {
  //     if c < '0' || c > '9' {
  //       apnd := strings.Replace(string(c), string(c), "", 1)
  //       s = append(s, apnd)
  //     }else{
  //       s = append(s, string(c))
  //     }
  //   }
  // }
  //
  // fmt.Println(s)

  k := true
  for _, c := range cf {
    if c < '0' || c > '9' {
      k = false
    }
  }
  if k == false {
		fmt.Println("error")
	}
}
