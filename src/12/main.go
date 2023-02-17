package main

import (
	"12/set"
	"fmt"
)

func main() {
	s := set.NewSet()
	s.Add("cat", "cat", "dog", "cat", "tree")
	v := s.Get()
	fmt.Println(v)
	s.Delete("cat", "dog")
	v = s.Get()
	fmt.Println(v)
}
