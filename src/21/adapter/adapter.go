package adapter

type SliceStorage struct {
	storage []any
}

func (s *SliceStorage) Send(v ...any) {
	for _, val := range v {
		s.storage = append(s.storage, val)
	}
}

func (s *SliceStorage) Get() []any {
	return s.storage
}

type sliceStorageAdapter struct { // Я хочу работать со слайсом вместо мапы, поэтому написал адаптер
	Storage *SliceStorage
}

func NewSliceAdapter() *sliceStorageAdapter {
	return &sliceStorageAdapter{Storage: &SliceStorage{storage: []any{}}}
}

func (s *sliceStorageAdapter) Send(v ...any) { // адаптируем метод Send(особо нечего адаптирвоать))))
	s.Storage.Send(v...)
}

func (s *sliceStorageAdapter) Get() map[any]any { // адптируем метод Get(), создаем слайс и записываем в него данные из мапы и возвращаем )
	res := make(map[any]any)
	for _, v := range s.Storage.Get() {
		res[v] = struct{}{}
	}
	return res
}
