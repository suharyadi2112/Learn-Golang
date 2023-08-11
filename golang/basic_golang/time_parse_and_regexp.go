package main
import (
  "fmt"
  "time"
  "regexp"
  "strings"
)
func main() {

  var date time.Time

  // date, _ = time.Parse("02-01-06", "06-26-78")
  date, _ = time.Parse("02/01/2006", "05/04/1996")

  fmt.Println(date, "change format")

 inputDate := "28 nov 2016"
 layOut := "2 Jan 2006"

 timeStamp, err := time.Parse(layOut, inputDate)

 fmt.Println("Test 1 date : ", inputDate)

 if err != nil {
         fmt.Println(err)
 } else {

         year, month, day := timeStamp.Date()
         fmt.Printf("Date : [%d]year : [%d]month : [%d]day \n", year, month, day)

         // mm-dd-yyyy
         fmt.Println("mm-dd-yyyy date format : ", timeStamp.Format("01-02-2006"))

 }

  fmt.Println("-------------TIME PARSE------------------")
  input := "06-26-78"
  layout := "01-02-06"
  t, _ := time.Parse(layout, input)
  fmt.Println(t)                       // 2017-08-31 00:00:00 +0000 UTC
  fmt.Println(t.Format("02-Jan-2006")) // 31-Aug-2017
  fmt.Println(t.Format("02/01/2006"))

  fmt.Println("-------------FIND AND REMOVE CHARACTER------------------")

  res := strings.ReplaceAll("5 Oktober 1991;", ";", "")
  fmt.Println(res)


  fmt.Println("--------REGEX---------")
  str1 := "31-07-2010"
	re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])-((19|20))")

	fmt.Printf("\nDate: %v :%v\n", str1, re.MatchString(str1))

  fmt.Println("-------------Cek Validasi Tanggal-----------------")

  inputt := "12-30-78"
	layoutt := "01-02-06"
	tt, errDate := time.Parse(layoutt, inputt)// parse tanggal
  if errDate != nil{
    fmt.Println(false)
    return
  }
  ress := tt.Format("02/01/2006") //mengembalikan format yang digunakan pada MappingBirthDate()

	fmt.Println(ress)

}
