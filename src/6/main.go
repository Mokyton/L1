package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	go Run1()
	go Run2()
	go Run3()
	go Run4()
	go Run5()
	time.Sleep(10 * time.Second)
}

// Run1 Завершение горутины при помощи канала отмены
func Run1() {
	cancel := make(chan struct{})
	defer close(cancel)
	go func() {
		defer fmt.Println("Exit with cancel channel")
		for {
			select {
			case <-cancel:
				return
			}
		}
	}()
	time.Sleep(1 * time.Second)
	cancel <- struct{}{}
	time.Sleep(1 * time.Second)
}

//Run2 Завершение горутины проверя с помощью проверки канала на закрытие
func Run2() {
	stream := make(chan int)
	go func() {
		defer fmt.Println("defer with closed channel")
		for {
			select {
			case _, ok := <-stream:
				if !ok {
					return
				}

			}
		}
	}()

	for i := 0; ; i++ {
		time.Sleep(1 * time.Second)
		stream <- i
		if i == 2 {
			close(stream)
			break
		}
	}
	time.Sleep(1 * time.Second)
}

// Run3 Завершение горутины с помощью контекста отмены
func Run3() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		defer fmt.Println("End with ctx with Cancel")
		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}()

	time.Sleep(3 * time.Second)
	cancelFunc()
	time.Sleep(1 * time.Second)
}

// Run4 Завершение горутины с помощью ctx with Deadline
func Run4() {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(4*time.Second))
	go func() {
		defer fmt.Println("End with ctx with DeadLine")
		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}()
	time.Sleep(5 * time.Second)
}

// Run5 завершение горутины с помощью ctx with Timeout
func Run5() {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	go func() {
		defer fmt.Println("End with ctx with Timeout")
		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}()
	time.Sleep(6 * time.Second)
}
