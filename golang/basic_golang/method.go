package main

import(
  "fmt"
  "strings"
)

type Student struct{
  Name string
  Age int
}

func(s Student) SayHello(){
  fmt.Println("hallo", s.Name)
}

func (s Student) getNameAt(i int) string {
    return strings.Split(s.Name, " ")[i-1]
}

func main() {
    var s1 = Student{"john wick", 21}
    s1.SayHello()

    var name = s1.getNameAt(2)
    fmt.Println("nama panggilan :", name)
}
