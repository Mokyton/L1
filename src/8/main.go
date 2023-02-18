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

	switch mask := int64(1) << i; val {
	case 0:
		*num = *num &^ mask
	case 1:
		*num = *num | mask
	}

}
