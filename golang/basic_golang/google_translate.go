package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://google-translate1.p.rapidapi.com/language/translate/v2/languages"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", "dc6ee37cfbmsh35729c62334db83p1d7ccdjsna78d6b6519b5")
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &res)
		
	if err != nil {
		fmt.Println("tes err")
	}

	fmt.Println(res)

	// fmt.Println(string(body))

}