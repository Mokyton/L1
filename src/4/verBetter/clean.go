package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"verNum1/workerpool"
)

func main() {
	var n int
	w := flag.Int("n", 5, "number of workers")
	flag.Parse()
	if *w == 0 {
		fmt.Println("Please enter numbers of workers")
		stdin := bufio.NewReader(os.Stdin)
		_, err := fmt.Fscan(stdin, &n)
		if err != nil {
			log.Fatalln("Invalid input")
		}

		*w = n
	}
	Run(*w)
}

func Run(n int) {
	writer := make(chan any)
	cancel := make(chan os.Signal)
	go Generator(writer, cancel)
	signal.Notify(cancel, os.Interrupt, syscall.SIGKILL, syscall.SIGTERM)
	var allTask []*workerpool.Task
	pool := workerpool.NewPool(allTask, n)
	go func() {
		defer fmt.Println("turn off the task generator")
		for {
			select {
			case _, ok := <-cancel:
				if !ok {
					pool.Stop()
					return
				}
				close(cancel)
				return
			default:
				time.Sleep(10 * time.Millisecond)
				task := workerpool.NewTask(func(workerID int, out chan any) {
					v, ok := <-out
					if !ok {
						return
					}
					fmt.Printf("Worker: %d, msg:%d\n", workerID, v)
				}, writer)
				pool.AddTask(task)
			}

		}
	}()
	pool.RunBackground()
}

func Generator(writer chan<- any, cancel chan os.Signal) {
	defer fmt.Println("turn off the msg generator")
	defer close(writer)
	for i := 0; ; i++ {
		select {
		case writer <- i:
		case _, ok := <-cancel:
			if !ok {
				return
			}
			close(cancel)
			return
		}
	}
}
