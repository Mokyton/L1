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
	var wg sync.WaitGroup
	wg.Add(3)
	in := make(chan any)
	doubleDat := make(chan any)
	go func() {
		writeToChan(data, in)
		wg.Done()
	}()
	go func() {
		doubleVal(in, doubleDat)
		wg.Done()
	}()

	go func() {
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
