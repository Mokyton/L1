package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
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

	pool := make(chan int, n)
	defer close(pool)
	for i := 1; i <= n; i++ {
		pool <- i
	}

	c := make(chan int)

	cancel := make(chan os.Signal)
	signal.Notify(cancel, os.Interrupt, syscall.SIGKILL, syscall.SIGTERM)

	go Generator(c, cancel)
	var exit bool
	for i := 0; ; i++ {
		select {
		case id := <-pool:
			go reader(pool, id, c)
		case _, ok := <-cancel:
			if !ok {
				exit = true
			} else {
				close(cancel)
				exit = true
			}
		}
		if exit {
			break
		}

	}
	for i := 0; i < n; i++ {
		<-pool
	}
}

func Generator(c chan<- int, cancel chan os.Signal) {
	defer close(c)
	for {
		select {
		case _, ok := <-cancel:
			if !ok {
				return
			}
			close(cancel)
			return
		case c <- rand.Intn(100):
		}
	}
}

func reader(worker chan<- int, workerID int, out <-chan int) {
	fmt.Printf("Worker: %d, msg:%d\n", workerID, <-out)
	worker <- workerID
}
