package workerpool

type Task struct {
	Data chan any
	f    func(workerID int, data chan any)
}

func NewTask(f func(workerID int, data chan any), data chan any) *Task {
	return &Task{f: f, Data: data}
}

func process(workerID int, task *Task) {
	task.f(workerID, task.Data)
}
