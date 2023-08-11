package main

import(
  "fmt"
  "time"
  "strings"
)

func main(){
  now := time.Now().UTC()
  nownotutc := time.Now().String()
  cekupdated_at := "2022-09-05 18:15:20.621" //1 jam kedepan sejak updated_at

	daridbsplit := strings.Split(cekupdated_at, ".")
  nownotutcsplit := strings.Split(nownotutc, ".")

  fmt.Println(nownotutcsplit[0])
  fmt.Println(daridbsplit[0])

  if nownotutcsplit[0] < daridbsplit[0] {
    fmt.Println("masih oke")
  }else{
    fmt.Println("kirim reminder")
  }


  startdate, enddate := GetDateByDays(now, 2)

  fmt.Println(startdate, enddate, "cek isi var")
}

func GetDateByDays(now time.Time, days int) (startDate, endDate time.Time) {
	startDate = now.AddDate(0, 0, -days-1)
	endDate = now.AddDate(0, 0, -days)

	return startDate, endDate
}
