package main

import(
	"fmt"
	"time"
)

 func main(){
	now := time.Now()

	fmt.Println(now)

	if now.Month() == time.January && now.Day() == 1 {
		fmt.Println("Happy New Year!")
	} else {
		fmt.Println("Happy New Year will come soon!")
	}

 }
