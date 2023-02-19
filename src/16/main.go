package main

import "fmt"

func main() {
	data := []int{3, 2, 1}
	fmt.Println(data)
	fmt.Println(quickSortStart([]int{3, 2, 1}))
	data = []int{18, 64, -25, 2, -7, 12}
	fmt.Println(data)
	Quicksort(data)
	fmt.Println(data)
}

// ver # 1

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high) // получаем индекс опорного элемент и измененный слайс
		arr = quickSort(arr, low, p-1)     // запускаем рекурсию для части слайса меньше опроного элемента
		arr = quickSort(arr, p+1, high)    // запускаем рекурсию для части слайса больше опроного элемента
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) { //
	pivot := arr[high] // опорный элемент
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot { // если элемент меньше опроного меняем его местами со следующим
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i] // делаем так, чтобы больший элемент был последним
	return arr, i                         // возращаем измененный слайс и индекс новго опроного элемента
}

// ver # 2

func Quicksort(ar []int) {
	if len(ar) <= 1 { // если размер слайса меньше или равен 1 сортировать не нужно (прекращение рекурсии)
		return
	}

	split := Partition(ar) // получаем индекс опорного элемента

	Quicksort(ar[:split]) // сортируем часть слайса до опорного элемента
	Quicksort(ar[split:]) // сортируем часть после опорного элемента с ним ключительно
}

func Partition(ar []int) int {
	pivot := ar[len(ar)/2] // выбирае опорный элемент берем из середины

	left := 0
	right := len(ar) - 1

	for {
		for ; ar[left] < pivot; left++ {
		}

		for ; ar[right] > pivot; right-- {
		}

		if left >= right { // возвращаем новый индекс опорного элемента
			return right
		}
		ar[left], ar[right] = ar[right], ar[left] // swap
	}
}
