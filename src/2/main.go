package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {

	data := []int{2, 4, 6, 8, 10}
	versionNumber1(data)
	versionNumber2(data)
	versionNumber3(data)
}

func versionNumber1(data []int) { // Примитив синхронизации WaitGroup
	stdout := bufio.NewWriter(os.Stdout)
	defer stdout.Flush()
	var wg sync.WaitGroup
	for i := 0; i < len(data); i++ {
		wg.Add(1)
		go func(i int) {
			data[i] *= data[i]
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 0; i < len(data); i++ {
		fmt.Fprintln(stdout, data[i])
	}

}

func versionNumber2(data []int) { // Синхронизация при помощи сигнального канала
	stdout := bufio.NewWriter(os.Stdout)
	defer stdout.Flush()
	done := make(chan struct{})
	for i := 0; i < len(data); i++ {
		go func(i int) {
			data[i] *= data[i]
			done <- struct{}{}
		}(i)
	}

	for i := 0; i < len(data); i++ {
		<-done
	}
	close(done)
	for i := 0; i < len(data); i++ {
		fmt.Fprintln(stdout, data[i])
	}
}

func versionNumber3(data []int) { // Таже синхронизация при помощи сигнального канала,
	stdout := bufio.NewWriter(os.Stdout) // только данные просто записываются сразу в stdout и slice не изменяется
	defer stdout.Flush()                 // использую mutex, чтобы небыло гонки при записи в stdout
	done := make(chan struct{})
	var lock sync.Mutex
	for i := 0; i < len(data); i++ {
		i := i
		go func() {
			lock.Lock()
			fmt.Fprintln(stdout, data[i]*data[i])
			lock.Unlock()
			done <- struct{}{}
		}()
	}

	for i := 0; i < len(data); i++ {
		<-done
	}
	close(done)
}
