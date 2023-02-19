package main

import (
	"21/adapter"
	"21/data/mapModel"
	"fmt"
)

func main() {
	mapStore := mapModel.New()
	adapterSliceStore := adapter.NewSliceAdapter()
	mapStore.Send(1, 2, 3, 4, 5)
	adapterSliceStore.Send([]any{6, 7, 8, 9, 10}...)
	fmt.Println(mapStore.Get())                  //  храниться в виде мапы
	fmt.Println(adapterSliceStore.Storage.Get()) // храниться в виде слайса
	client := client{}
	client.ShowData(mapStore)
	client.ShowData(adapterSliceStore) // адаптация
}
