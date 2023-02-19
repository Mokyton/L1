package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

var justString string // использование глобальной переменной bad practice

// Минусы при создании строки выделяется слишком большое количество памяти
// Размер строки выделяется в байтах, срез из 100 элементов стркоки может быть не коректным при работе с unicode,
// так как некоторые символы могут весить больше 1 байта

func someFunc() {
	v := createHugeString(1 << 10) // 1024
	justString = v[:100]           // есть вероятность словить панику out of range
}

func createHugeString(size int) string {
	res := make([]byte, size)

	for i := 0; i < size; i++ {
		res[i] = byte(rand.Intn(1103-1040) + 1040)
	}

	return string(res)
}

func main() {
	//someFunc()
	mySomeFunc()
}

// My version

func mySomeFunc() {
	myJustString := string(make([]byte, 100))
	v := createHugeString(1 << 10)
	if len(v)-1 >= 100 {
		myJustString = v[:100]
	} else {
		log.Fatal(errors.New("out of range"))
	}
	fmt.Println(myJustString)
}
