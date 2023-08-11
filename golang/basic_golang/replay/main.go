package main

import (
    "fmt"
    "strings"
    "net/http"
    "io/ioutil"
)

func main() {
    url := "https://api.twitter.com/2/tweets"

    payload := strings.NewReader(`{
        "text": "Hello"
    }`)

    req, err := http.NewRequest("POST", url, payload)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    req.Header.Add("cookie", "guest_id_marketing=v1%253A168775179625857235; guest_id_ads=v1%253A168775179625857235; personalization_id=%22v1_mqMUly7T7DQPWVbv8xR0Lg%3D%3D%22; guest_id=v1%253A168775179625857235")
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Authorization", `OAuth oauth_consumer_key="S5DZlBmM6w6a9eb77FehZcvXe", oauth_nonce="KGfz7qhkTBi5rc744CrF7dWBcR1IzOaV", oauth_signature="f1uSs%2BbbMqPIeORM9hP2MxpyQYw%3D", oauth_signature_method="HMAC-SHA1", oauth_timestamp="1687753665", oauth_token="1302023382964215809-OI3sQWraHXQtZLKZUCu19r1mgift0c", oauth_version="1.0"`)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    fmt.Println(res)
    fmt.Println(string(body))
}


// package main

// import (
//     "fmt"
//     "log"
//     // "os"

//     "github.com/ChimeraCoder/anaconda"
// )

// func main() {
//     // consumerKey := os.Getenv("S5DZlBmM6w6a9eb77FehZcvXe")
//     // consumerSecret := os.Getenv("6heDDW598nbTS8nfZTJE7jyWtmPkiuIXx39J8LrW7gEAeUsqbN")
//     // accessToken := os.Getenv("1302023382964215809-OI3sQWraHXQtZLKZUCu19r1mgift0c")
//     // accessTokenSecret := os.Getenv("GNk8T3QQJuthZKW172ViPl3WSUGE3zHAPNgDu7nBb9rY8")

//     // // Twitter API authentication
//     // anaconda.SetConsumerKey(consumerKey)
//     // anaconda.SetConsumerSecret(consumerSecret)

//     api := anaconda.NewTwitterApiWithCredentials("1302023382964215809-OI3sQWraHXQtZLKZUCu19r1mgift0c", "GNk8T3QQJuthZKW172ViPl3WSUGE3zHAPNgDu7nBb9rY8", "S5DZlBmM6w6a9eb77FehZcvXe", "6heDDW598nbTS8nfZTJE7jyWtmPkiuIXx39J8LrW7gEAeUsqbN")

//     fmt.Println(api)

//     // api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

//       // ...

//     // searchParams := anaconda.NewSearchTweetParams("#indonesia")
//     // searchParams.Count = 2 // Set the number of tweets to retrieve

//     // fmt.Println(searchParams, "tes-")

//     // Search for tweets
//     // searchQuery := "#sapi"
//     // results, err := api.GetSearch(searchQuery, nil)
//     // if err != nil {
//     //     log.Fatal(err)
//     // }

//     // fmt.Println(results)

//     // Iterate through the tweets and reply to each one
//     // for _, tweet := range results.Statuses {
//     //     replyText := fmt.Sprintf("@%s Your reply message here!", tweet.User.ScreenName)
//     //     _, err := api.PostTweet(replyText, nil)
//     //     if err != nil {
//     //         log.Println("Error replying to tweet:", err)
//     //     } else {
//     //         log.Println("Replied to tweet successfully!")
//     //     }
//     // }

//     //Post Tweet
//     tweetText := "Hello, world! This is my tweet."
//     //api := createTwitterClient()



//     tweet, err := api.PostTweet(tweetText, nil)
//     if err != nil {
//         log.Fatal(err)
//     }

//     log.Printf("Tweet posted successfully! Tweet ID: %d\n", tweet.Id)
// }
