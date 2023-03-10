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

// Run использую context.WithTimeout через указанное время dur ctx.Done() вернет канал прочитыва из него мы остановим generator
func Run(n int) {
	writer := make(chan int)
	dur := time.Second * time.Duration(n)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(dur))
	fmt.Println("start")
	go Generator(ctx, writer)
	for {
		select {
		case v, ok := <-writer:
			if !ok { // когда функция Generator завершится в ней закроется канал из которого мы читаем.
				fmt.Println("END") // если канал закрыт мы выходим из бесконечного цикла
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
