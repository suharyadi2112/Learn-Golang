package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://api.twitter.com/2/tweets/search/all?query=tes"
  method := "GET"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  req.Header.Add("cookie", "guest_id_marketing=v1%253A168775179625857235; guest_id_ads=v1%253A168775179625857235; personalization_id=%22v1_mqMUly7T7DQPWVbv8xR0Lg%3D%3D%22; guest_id=v1%253A168775179625857235")
  req.Header.Add("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAAOK1oQEAAAAAk2nTi%2B8c4z4igs1CAphsVBKisqA%3D0flTUwkB4s3cTg3yZ9KKptzkRDOnchR9Nr3Ht3iWF1SKDiCjPN")

  if err != nil {
    fmt.Println(err)
    return
  }
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}
