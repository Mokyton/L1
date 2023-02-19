package mapModel

type MapStorage struct {
	Storage map[any]any
}

func (m *MapStorage) Send(v ...any) { //  Запись данных в мапу
	for _, val := range v {
		m.Storage[val] = struct{}{}
	}
}

func (m *MapStorage) Get() map[any]any { // Возвращает данные ввиде мапы
	return m.Storage
}

func New() *MapStorage { // конструктор
	return &MapStorage{Storage: make(map[any]any)}
}
