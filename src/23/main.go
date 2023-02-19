package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(Delete([]int{1, 2, 3, 4, 5, 6, 7}, 2))
	fmt.Println(Delete([]int{1, 2, 3, 4, 5, 6, 7}, 5))
	fmt.Println(Delete([]int{1, 2, 3, 4, 5, 6, 7}, -1))
	fmt.Println(Delete([]int{1, 2, 3, 4, 5, 6, 7}, 8))
}

func Delete(src []int, i int) ([]int, error) {
	if i < 0 || i > len(src)-1 { // Проверяем, что такой индекс существует
		return nil, errors.New(fmt.Sprintf("Error: invalid index %d", i))
	}
	cur := src[:i]                          // длеаем срез всех элементов до удаляемого
	cur = append(cur, src[i+1:len(src)]...) // добавляем все элементы, которые идут после удаляемого
	return cur, nil
}
