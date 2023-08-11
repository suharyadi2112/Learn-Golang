package main
import (
  "fmt"
  "reflect"// reflect digunakan untuk melihat tipe data dari sebuah variable
)

type Student struct {
  Name string
  Grade int
}

func main() {
  var number = 11
  var numberreflect = reflect.ValueOf(number)//cek tipe data number

  fmt.Println("tipe  variabel :", numberreflect.Type())//cek tipe data number
  fmt.Println("tipe  variabel :", numberreflect.Interface())//cek value data number

  var nilai = numberreflect.Interface().(int)//cek value data number
  fmt.Println(nilai)

  //----------------------------//
  //reflect mengambil dari struct
  // Pengaksesan Informasi Property Variabel Objek
  var cekHasil = &Student{Name : "suharyadi", Grade : 90}
  cekHasil.CekStudent()
}

// Pengaksesan Informasi Property Variabel Objek
func (s *Student) CekStudent()  {

  //dapatkan struct dan isi dari penjabaran struct Student
  var hasilStudent = reflect.ValueOf(s)

  if hasilStudent.Kind() == reflect.Ptr {//pengecekan apakah variabel objek tersebut merupakan pointer atau tidak
      //jika bernilai true maka variabel adalah pointer). jika ternyata variabel memang pointer, maka perlu diambil objek reflect aslinya dengan cara memanggil method
      hasilStudent = hasilStudent.Elem()
  }

  var cekStudent = hasilStudent.Type()

  for i := 0; i < hasilStudent.NumField(); i++ {//hitung total struct Student menggunakan NumField()
      fmt.Println("nama      :", cekStudent.Field(i).Name)//akan mengembalikan nama property
      fmt.Println("tipe data :", cekStudent.Field(i).Type)//mengembalikan tipe data property
      fmt.Println("nilai     :", hasilStudent.Field(i).Interface())//mengembalikan nilai property dalam bentuk interface{}
      fmt.Println("")
  }
}
