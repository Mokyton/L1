package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type verNum1 struct {
	sync.Mutex
	counter int
}

type verNum2 struct {
	counter int32
}

func main() {
	Run1()
	Run2()
}

func Run1() int {
	var wg sync.WaitGroup
	counter := verNum1{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Lock()
			counter.counter++
			counter.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter.counter)
	return counter.counter
}

func Run2() int {
	var wg sync.WaitGroup
	counter := verNum2{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter.counter, 1)
		}()
	}
	wg.Wait()
	fmt.Println(counter.counter)
	return int(counter.counter)
}
