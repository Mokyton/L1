package main

import (
	"fmt"
	"sync"
)

type MapConcurrency struct {
	storage map[any]any
	sync.RWMutex
}

func main() {
	m := NewMapConcurrency()
	test := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4,
		"e": 5, "dsad": 50, "cxc": 89, "das": 1, "sda": 54, "bv": 2}
	var wg sync.WaitGroup

	for k, v := range test {
		wg.Add(1)
		go func(k, v any) {
			defer wg.Done()
			m.Add(k, v)
		}(k, v)
	}
	wg.Wait()
	fmt.Println("Data len:", len(test), "Map len:", len(m.storage))
	if len(test) != len(m.storage) {
		fmt.Println(len(test), m.storage)
		fmt.Println("Error")
	}
	m.PrintMap()
}

func NewMapConcurrency() *MapConcurrency {
	m := make(map[any]any)
	return &MapConcurrency{storage: m}
}

func (m *MapConcurrency) Add(k any, v any) { // Записываем элементы в map
	m.Lock()
	m.storage[k] = v
	m.Unlock()
}

func (m *MapConcurrency) Delete(k any) { // Удаляем элементы из map
	m.Lock()
	delete(m.storage, k)
	m.Unlock()
}

func (m *MapConcurrency) GetValue(k any) any { // Получаем Value из map
	m.RLock()
	defer m.RUnlock()
	return m.storage[k]
}

func (m *MapConcurrency) PrintMap() { // принтуем map
	m.RLock()
	fmt.Println(m.storage)
	m.RUnlock()
}
