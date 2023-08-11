package main

import (
  "net/http"
)

const USERNAME = "batman"
const PASSWORD = "secret"

// fungsi Auth() adalah memvalidasi apakah request merupakan valid basic auth request, dan juga apakah credentials yang dikirim cocok dengan data pada aplikasi kita.
func Auth(w http.ResponseWriter, r *http.Request) bool {
    //Fungsi r.BasicAuth() mengembalikan 3 informasi: Username, Password dan Nilai balik ke-3 ini adalah representasi valid tidak nya basic auth request yang sedang berlangsung
    username, password, ok := r.BasicAuth()
    if !ok {
        w.Write([]byte(`something went wrong`))
        return false
    }
    //cek username dan password apakah cocok
    isValid := (username == USERNAME) && (password == PASSWORD)
    if !isValid {
        w.Write([]byte(`wrong username/password`))
        return false
    }
    return true
}
// Fungsi ini bertugas untuk memastikan bahwa request yang diperbolehkan hanya yang ber-method GET.
func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
    if r.Method != "GET" {
        w.Write([]byte("Only GET is allowed"))
        return false
    }

    return true
}
