package workerpool

type Task struct {
	Data chan any                          // канал для чтения
	f    func(workerID int, data chan any) // функция для чтения
}

func NewTask(f func(workerID int, data chan any), data chan any) *Task { // конструктор Задач
	return &Task{f: f, Data: data}
}

func process(workerID int, task *Task) { // вызыво задачи
	task.f(workerID, task.Data)
}
