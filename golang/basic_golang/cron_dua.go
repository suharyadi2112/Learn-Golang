package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func main() {
    // Run the function "CallText" every minute
    // gocron.Every(1).Second().Do(CallText)
    gocron.Every(0).Day().At().Do(CallText)

    // Start the scheduler
    <- gocron.Start()
}

func CallText() {
    fmt.Println("This Text")
}