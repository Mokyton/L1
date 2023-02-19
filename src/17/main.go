package main

import (
	"errors"
	"fmt"
	"sort"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index, _ := binSearch(data, 10)
	fmt.Println(index)
}

func binSearch(data []int, val int) (int, error) {
	if !sort.SliceIsSorted(data, func(i, j int) bool { // валидация на поданные значения
		return data[i] < data[j]
	}) {
		return -1, errors.New("Error: unsorted slice ")
	}

	left := 0              //  стандартный алгоритм из Грокаем алгоритмы
	right := len(data) - 1 // проверяем центральный элемент если он нам походим возвращаем его индекс
	for left <= right {    // если центральный элемент меньше ожидаемого нижняя граница поиска становится индекс цетрального элемента +1
		mid := (right + left) / 2 // если центральный элемент больше верхняя граница поиска становится центральный индекс - 1
		if val == data[mid] {
			return mid, nil
		}
		if val > data[mid] {
			left = mid + 1
		}
		if val < data[mid] {
			right = mid - 1
		}
	}
	return -1, errors.New("This elements doesn't exist ")
}
