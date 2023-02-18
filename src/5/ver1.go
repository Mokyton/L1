package main

import (
	"bufio"
	"context"
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
	Run(*w)
}

func Run(n int) {
	writer := make(chan int)
	dur := time.Second * time.Duration(n)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(dur))
	fmt.Println("start")
	go Generator(ctx, writer)
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

func Generator(ctx context.Context, in chan<- int) {
	defer close(in)
	for {
		time.Sleep(100 * time.Millisecond)
		select {
		case <-ctx.Done():
			fmt.Println("Generator stopped")
			return
		case in <- rand.Intn(100):

		}
	}
}
