package data

type Storage interface { // интерфейс для работы с сервисом
	Send(v ...any)
	Get() map[any]any
}
