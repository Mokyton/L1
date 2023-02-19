package main

import (
	"fmt"
)

func main() {
	fmt.Println("欢迎!", "--", reverseString("欢迎!"))
	fmt.Println("Hello", "--", reverseString("Hello"))
}

func reverseString(word string) string {
	n := len(word)
	reversed := make([]rune, n)
	for i, char := range word { //  использую range, потому что он возвращает rune
		reversed[n-i-1] = char
	}
	return string(reversed)
}
