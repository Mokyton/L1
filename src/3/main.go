package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	data := []int{2, 4, 6, 8, 10}
	fmt.Println(versionNumber1(data...))
	fmt.Println(versionNumber2(data...))

}

func versionNumber1(src ...int) int { // подсчет с использованием Mutex и Wait Group
	var wg sync.WaitGroup
	var lock sync.Mutex
	var sum int
	for _, v := range src {
		v := v
		wg.Add(1)
		go func() {
			lock.Lock()
			sum += v * v
			lock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return sum
}

func versionNumber2(v ...int) int { // подсчет с использованием atomics
	var sum atomic.Int32
	done := make(chan struct{})
	for _, v := range v {
		go func(v int) {
			sum.Add(int32(v * v))
			done <- struct{}{}
		}(v)
	}
	for i := 0; i < len(v); i++ {
		<-done
	}
	close(done)
	return int(sum.Load())
}
