package main

import "fmt"

func main() {
	data := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	Run(data)
}

func round(num float32) int {
	return int(num) / 10 * 10
}

func Run(data []float32) {
	storage := make(map[int][]float32)
	for i := 0; i < len(data); i++ {
		key := round(data[i])
		storage[key] = append(storage[key], data[i])
	}
	fmt.Println(storage)
}
