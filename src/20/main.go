package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("snow dog sun", "--", reversWordsInSlice("snow dog sun"))
	fmt.Println("我不会 说汉语", "--", reversWordsInSlice("我不会 说汉语"))
}

func reversWordsInSlice(src string) string {
	buf := strings.Fields(src)
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i] // swap первого и последнего элемента слайса 0 9, 1 8, 2 7, 3 6, 4 5, 5 > 4 stop
	}
	return strings.Join(buf, " ")
}
