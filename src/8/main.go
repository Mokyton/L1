package main

import "fmt"

func main() {
	num := int64(1)
	setBit(&num, 2, 1)
	fmt.Println(num)
}

func setBit(num *int64, i, val int64) {
	if (val < 0 && val > 1) || (i > 63 && i < 0) { // Валидация на крайние значения входных данных
		return
	}

	switch mask := int64(1) << i; val { // создаем маску и на основе значения применяем ее
	case 0:
		*num = *num &^ mask // побитовые "и не" зануляем бит 101 &^ (1 << 2) = 101 &^ 100 = 001
	case 1:
		*num = *num | mask // побитовая или если хотя бы одно из значений true(1) ставим 1 (по дефолту в маске 1) 1 | 0 = 1
	}

}
