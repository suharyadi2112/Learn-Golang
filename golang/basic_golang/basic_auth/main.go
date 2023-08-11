package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//RUNNING GO RUN HARUS KETIGA FILE DI DALAM BASIC_AUTH KARENA FILE TERPISAH DAN TIDAK DI IMPORT
//-----------------HTTP Basic Authentication---------------//
func main() {

    //Routing
    http.HandleFunc("/student", ActionStudent)

    server := new(http.Server)
    server.Addr = ":9000"// webserver running

    fmt.Println("server started at localhost:9000")//running at port local
    server.ListenAndServe()//eksekusi webserver local
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
    //Validasi !Auth(w, r); untuk mengecek apakah request merupakan valid basic auth request atau tidak.
    if !Auth(w, r){
      return
    }
    // Validasi !AllowOnlyGET(w, r); gunanya untuk memastikan hanya request dengan method GET yang diperbolehkan masuk.
    if !AllowOnlyGET(w, r) {
      return
    }

    //cek request ini memiliki parameter student id atau tidak
    if id := r.URL.Query().Get("id"); id != "" {
        OutputJSON(w, SelectStudent(id))//jika ada id dalam request
        return
    }
    //jika tidak ada id request.
    OutputJSON(w, GetStudents())
}
func OutputJSON(w http.ResponseWriter, o interface{}) {
    // Fungsi ini digunakan untuk mengkonversi data menjadi JSON string.
    res, err := json.Marshal(o)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(res)
}
