package main

import (
	"fmt"
	// "math/rand"
	// "time"
	// "errors"

)

func main() {
	// // Set the seed for the random number generator
	// rand.Seed(time.Now().UnixNano())

	// // Define the characters that will be included in the password
	// chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()")

	// // Generate a random password of length 10
	// password := make([]rune, 10)
	// for i := range password {
	// 	password[i] = chars[rand.Intn(len(chars))]
	// }

	// // Print the generated password
	// fmt.Println(string(password))

	password := "@Tes123"

	counterLower := 0
	counterUpper := 0
	counterNumber := 0
	counterSymbol := 0
	for _, p := range password {
		if p >= 'a' && p <= 'z' {
			counterLower = 1
		} else if p >= 'A' && p <= 'Z' {
			counterUpper = 1
		} else if p >= '1' && p <= '9' {
			counterNumber = 1
		} else {
			counterSymbol = 1
		}
	}

	total := counterLower + counterUpper + counterNumber + counterSymbol
	if total < 3 {
		fmt.Println("sf")
	}

	fmt.Println("tes")
}