package main

import (
	"21/data"
	"fmt"
)

type client struct {
}

func (c *client) ShowData(storage data.Storage) {
	fmt.Println(storage.Get())
}
