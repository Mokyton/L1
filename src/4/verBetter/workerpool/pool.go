package workerpool

type Pool struct {
	Tasks   []*Task
	Workers []*Worker

	Size          int
	collector     chan *Task
	runBackground chan struct{}
}

func (p *Pool) AddTask(task *Task) { // добавление задачек
	p.collector <- task
}

func NewPool(tasks []*Task, size int) *Pool { // конструктор бассейна
	return &Pool{
		Tasks:     tasks,
		Size:      size,
		collector: make(chan *Task, 1000),
	}
}

func (p *Pool) RunBackground() {

	for i := 1; i <= p.Size; i++ {
		worker := NewWorker(p.collector, i)
		p.Workers = append(p.Workers, worker)
		go worker.StartBackground()
	}

	for i := range p.Tasks { // закидываем задачки, которые подготовили заранее
		p.collector <- p.Tasks[i]
	}

	p.runBackground = make(chan struct{})
	<-p.runBackground // выжидаем пока все workers закончат работать
}

func (p *Pool) Stop() {
	for i := range p.Workers {
		p.Workers[i].Stop()
	}
	p.runBackground <- struct{}{}
}
