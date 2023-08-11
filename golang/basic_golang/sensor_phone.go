package main

import(
	"strings"
	"fmt"
	"math"
)


func main(){
	fmt.Println(CensorPhoneFormat("08228856855", 0.6, 2))
}

func CensorPhoneFormat(name string, censorFactor float64, firstDigit int) string {
	if len(name) < firstDigit+1 {
		return name
	}

	first := name[0:firstDigit]
	fmt.Println(first)
	name = name[firstDigit:len(name)]
	words := strings.Fields(name)
	for i, s := range words {
		if len(s) > 0 {
			toCensor := int(math.Floor(float64(len(s)) * censorFactor))
			s = strings.Repeat("*", toCensor) + "-" + s[toCensor:len(s)]
		}
		words[i] = s
	}

	return "+" + first + "-" + strings.Join(words, " ")
}