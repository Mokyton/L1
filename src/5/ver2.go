package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	w := flag.Int("d", 5, "program duration in seconds")
	flag.Parse()
	if *w == 0 {
		var n int
		fmt.Println("Please enter program duration in seconds")
		stdin := bufio.NewReader(os.Stdin)
		_, err := fmt.Fscan(stdin, &n)
		if err != nil {
			log.Fatalln("Invalid input")
		}

		*w = n
	}
	Run1(*w)
}

func Run1(dur int) {
	writer := make(chan int)
	cancel := make(chan struct{})
	fmt.Println("start")
	go func() {
		time.Sleep(time.Duration(dur) * time.Second)
		cancel <- struct{}{}
	}()
	go Generator1(cancel, writer)
	for {
		select {
		case v, ok := <-writer:
			if !ok {
				fmt.Println("END")
				return
			}
			fmt.Println(v)
		}
	}
}

func Generator1(cancel chan struct{}, in chan<- int) {
	defer close(in)
	for {
		time.Sleep(100 * time.Millisecond)
		select {
		case <-cancel:
			fmt.Println("Generator stopped")
			return
		case in <- rand.Intn(100):

		}
	}
}
