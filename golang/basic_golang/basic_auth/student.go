package main
// import (
//   "fmt"
//   "encoding/json"
// )
var StudentData = []*Student{}

type Student struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Grade int32  `json:"grade"`
}

//fungsi ini mengembalikan semua data student
func GetStudents() []*Student {
    return StudentData
}

// Dan buat juga fungsi SelectStudent(id), fungsi ini mengembalikan data student sesuai dengan id terpilih.
func SelectStudent(id string) *Student {
    for _, each := range StudentData   {
        if each.Id == id {
            return each
        }
    }
    return nil
}
// func main jika ingin nampilkan data langsung dari main student
// func main(){
//   res, _ := json.Marshal(GetStudents())//data di marshal terlebih dahulu
//   var CekStructStudent []Student //data ditampung di dalam struct
//   json.Unmarshal(res, &CekStructStudent)// Unmarshal lalu tampung diStruct
//
//   //looping data menggunakan for
//   for lh := 0; lh < len(CekStructStudent); lh++ {
//       fmt.Println(CekStructStudent[lh].Name)
//   }
//
// }

// func init(), buat beberapa dummy data untuk ditampung pada variabel students.
// func init() bawaan dari golang
func init() {
    StudentData = append(StudentData, &Student{Id: "s001", Name: "bourne", Grade: 2})
    StudentData = append(StudentData, &Student{Id: "s002", Name: "ethan", Grade: 2})
    StudentData = append(StudentData, &Student{Id: "s003", Name: "wick", Grade: 3})
}
