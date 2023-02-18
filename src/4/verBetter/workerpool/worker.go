package workerpool

import (
	"fmt"
)

type Worker struct {
	ID       int
	taskChan chan *Task
	quit     chan struct{}
}

func NewWorker(channel chan *Task, ID int) *Worker { // Конструктор работяг
	return &Worker{
		ID:       ID,
		taskChan: channel,
		quit:     make(chan struct{}), // канал для завершения
	}
}

func (wr *Worker) StartBackground() { // запуск работяги
	fmt.Printf("Starting worker %d\n", wr.ID)

	for {
		select {
		case task := <-wr.taskChan:
			process(wr.ID, task)
		case <-wr.quit:
			return
		}
	}
}

func (wr *Worker) Stop() { //  функция для отправки работяги домой
	fmt.Printf("Closing worker %d\n", wr.ID)
	go func() {
		wr.quit <- struct{}{}
	}()
}
