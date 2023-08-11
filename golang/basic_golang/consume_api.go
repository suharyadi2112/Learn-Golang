package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todos struct {
	UserID    int    `json:"userId"`
	Id 		  int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// Make a GET request to the API, here we will use
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/2")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return
	}
	// Defer closing the response body
	defer response.Body.Close()	
	// Read the response body
	data, _ := ioutil.ReadAll(response.Body)

	// Unmarshal the JSON data into a Todos struct
	var todos Todos

	// Decode JSON to Variabel Objek Struct
	json.Unmarshal(data, &todos)

	// Print the todo's with object, 
	fmt.Println("\n","title :",todos.Title,"\n","User ID :",todos.UserID,"\n", "ID :",todos.Id,"\n","Completed :",todos.Completed)
}