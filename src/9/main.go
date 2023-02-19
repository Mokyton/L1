package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	Run(data)
}

func Run(data []int) {
	var wg sync.WaitGroup // способ синхронизации WaitGroup
	wg.Add(3)
	in := make(chan any)        // Канал для записи чисел из слайса
	doubleDat := make(chan any) // Канал для записи квадратов чисел
	go func() {                 // читаем числа из слайса и записываем в канал
		writeToChan(data, in)
		wg.Done()
	}()
	go func() { // читаем числа из канала in и считаем их квадрат и записываем в канал doubleDat
		doubleVal(in, doubleDat)
		wg.Done()
	}()

	go func() { // читаем числа канала и пишем их в stdout
		Print(doubleDat)
		wg.Done()
	}()
	wg.Wait()
}

func writeToChan(data []int, in chan<- any) {
	defer close(in)
	for _, v := range data {
		in <- v
	}
}

func doubleVal(read <-chan any, write chan<- any) {
	defer close(write)
	for {
		select {
		case v, ok := <-read:
			if !ok {
				return
			}
			write <- v
		}
	}
}

func Print(read <-chan any) {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for {
		select {
		case v, ok := <-read:
			if !ok {
				return
			}
			fmt.Fprintln(out, v)
		}
	}
}
