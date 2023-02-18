package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
)

type Task interface {
	Execute(in <-chan int)
}

type Reader struct{}

func (r *Reader) Execute(in <-chan int) {
	fmt.Println(<-in)
}

type Pool struct {
	mu    sync.Mutex
	size  int
	tasks chan Task
	kill  chan os.Signal
	wg    sync.WaitGroup
}

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
	pool := NewPool(n)

	socket := make(chan int)
	t := Reader{}
	pool.Exec(t.Execute(socket))
	//msg := make(chan int)

	//pool.Resize(15)
	//for i := 0; i < 20; i++ {
	//	pool.Exec(ExampleTask(fmt.Sprintf("additional_%d", i+1)))
	//}

	defer pool.Close()

	pool.Wait()
}

func NewPool(size int, cancel chan os.Signal) *Pool {
	pool := &Pool{
		tasks: make(chan Task, 1024),
		kill:  make(chan struct{}),
	}
	pool.Resize(size)
	return pool
}

func (p *Pool) Resize(n int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for p.size < n {
		p.size++
		p.wg.Add(1)
		go p.worker()
	}
	for p.size > n {
		p.size--
		p.kill <- struct{}{}
	}
}

func (p *Pool) worker() {
	defer p.wg.Done()
	for {
		select {
		case task, ok := <-p.tasks:
			if !ok {
				return
			}
			task.Execute()
		case <-p.kill:
			return
		}
	}
}

func (p *Pool) Close() {
	close(p.tasks)
}

func (p *Pool) Wait() {
	p.wg.Wait()
}

func (p *Pool) Exec(task Task) {
	p.tasks <- task
}

func Generator(writer chan<- int, cancel chan struct{}) {
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
