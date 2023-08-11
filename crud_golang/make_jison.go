package main
 
import (
    "fmt" 
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "io/ioutil"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/lib/pq"
    // "os"
)

type Article struct {
    Id string `json:"id"`
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

type User struct {

    Id int `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`

}

type JsonResponse struct {
    Status   string
    Data    []User `json:"data"`
}

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "12345678"
  dbname   = "Golearn"
)

var Articles []Article
var Users []User

func connect() (*sql.DB, error) {

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, err
    }

    return db, nil
}

func AllUser(w http.ResponseWriter, r *http.Request){

    db, err := connect()//koneksi
    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json")

    if err != nil {
        fmt.Println(err.Error()) 
    }

    defer db.Close()

    row, err := db.Query("Select * from users")//query dari mysql
    if err != nil {
        fmt.Println(err.Error())
        return 
    }

    defer db.Close()

    var result []User

    for row.Next(){
        var each = User{}//User dari struct
        var err = row.Scan(&each.Id, &each.Name, &each.Email)

        if err != nil {
            fmt.Println(err.Error())
            return 
        }
        result = append(result, each)
    }

    if err = row.Err(); err != nil {
        fmt.Println(err.Error())
        return
    }

    var response = JsonResponse{Status: "success", Data: result}

    json.NewEncoder(w).Encode(response)

    fmt.Println("Endpoint hit Alluser")

    return
}

func NewUser(w http.ResponseWriter, r *http.Request){

    db, err := connect()//koneksi
    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json")

    if err != nil {
        fmt.Println(err.Error()) 
    }

    defer db.Close()

    // body, _ := ioutil.ReadAll(r.Body)
    // keyVal := make(map[string]string)
    // json.Unmarshal(body, &keyVal)
    //cara get key dan value post dari json string, bukan dari form value
    // name := keyVal["name"]  
    // email := keyVal["email"]

    //gunakan fungsi ParseForm() jika imputan berasal dari form 
    name :=  r.FormValue("name")
    email :=  r.FormValue("email")

    sqlStatement := `INSERT INTO users (name, email) VALUES ($1, $2)`

    _, err = db.Exec(sqlStatement, name, email)
    if err != nil {
      panic(err)
    }

    defer db.Close()

    var response = JsonResponse{Status: "success"}

    json.NewEncoder(w).Encode(response)

}

func DelUser(w http.ResponseWriter, r *http.Request){
    
    db, err := connect()//koneksi
    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json")

    if err != nil {
        fmt.Println(err.Error())
        return 
    }

    vars := mux.Vars(r)//mux digunakan untuk get id
    id := vars["id"]//get id menggunakan mux berdasarkan parameter dilink

    sqlQuery := `delete FROM users where id = $1`
     _, err = db.Exec(sqlQuery, id)

     if err != nil {
        fmt.Println(err.Error())
        return 
     }

    fmt.Println("EndPoint hit Deluser")
}

func PutUser(w http.ResponseWriter, r *http.Request){

    // db, err := connect()
    //tes 

    vars := mux.Vars(r)//mux digunakan untuk get id
    id := vars["id"]

    name :=  r.FormValue("name")
    email :=  r.FormValue("email")

    cek := map[string]string{
        "name" : name,
        "email" : email,
        "id" : id,
    }

    json.NewEncoder(w).Encode(cek)

    // if err != nil {
    //     json.NewEncoder(w).Encode(err.Error())
    //     return 
    // }
}


//mengembalikan semua data dari artikel
func returnAllArticles(w http.ResponseWriter, r *http.Request){

    fmt.Println("Endpoint Hit: returnAllArticles")

    for _, article := range Articles {
        json.NewEncoder(w).Encode(article)
    }

    // json.NewEncoder(w).Encode(Articles)
}   

//mengembalikan salah 1 artikel berdasarkan ID
func returnSingleArticle(w http.ResponseWriter, r *http.Request){

    vars := mux.Vars(r)
    var key string = vars["id"]

    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)        
        }
    }
    
    // fmt.Fprintf(w, "Key: " + key)

}
//create new article
func createNewArticle(w http.ResponseWriter, r *http.Request) {

    // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.    
    reqBody, _ := ioutil.ReadAll(r.Body)

    var article Article 
    json.Unmarshal(reqBody, &article)
    // update our global Articles array to include
    // our new Article
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)

}

//Hapus article
func deleteArticle(w http.ResponseWriter, r *http.Request){

    vars := mux.Vars(r)
    var key string = vars["id"]

    for index, article := range Articles {
        if article.Id == key {
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }

}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {

    //Routing 
    myRouter := mux.NewRouter().StrictSlash(true)

    //Solo requets melalui HTTP request
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/all", returnAllArticles)
    myRouter.HandleFunc("/SingleArticle/{id}", returnSingleArticle)
    myRouter.HandleFunc("/NewArticle", createNewArticle).Methods("POST")
    myRouter.HandleFunc("/DelArticle/{id}", deleteArticle).Methods("DELETE")
    myRouter.HandleFunc("/UpArticle/{id}", deleteArticle).Methods("PUT")

    //user
    myRouter.HandleFunc("/alluser", AllUser)
    myRouter.HandleFunc("/newuser", NewUser).Methods("POST")
    myRouter.HandleFunc("/deluser/{id}", DelUser).Methods("DELETE")
    myRouter.HandleFunc("/putuser/{id}", PutUser).Methods("PUT")

    fmt.Println("starting web server at http://localhost:9999/")
    log.Fatal(http.ListenAndServe(":9999", myRouter))
}

func main() {

    Articles = []Article{
        Article{Id: "1", Title: "Hello 1", Desc: "Deksripsi Satu", Content: "Article Content Satu"},
        Article{Id: "2", Title: "Hello 2", Desc: "Deskripsi Dua", Content: "Article Content Dua"},
        Article{Id: "3", Title: "Hello 3", Desc: "Deskripsi Tiga", Content: "Article Content Tiga"},
    }

    handleRequests()
}

