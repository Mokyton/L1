package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(uniqString("abcd"))
	fmt.Println(uniqString("aDcd"))
}

func uniqString(src string) bool {
	src = strings.ToLower(src) // Переводим все в нижний регистр
	set := make(map[rune]struct{})
	for _, v := range src { // получаем руну
		_, ok := set[v]
		if ok { //  если руна уже была записана значит в строке есть дубликаты
			return false
		}
		set[v] = struct{}{} // пустая структурка для экономии памяти =)
	}
	return true
}
