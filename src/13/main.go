package main

import "fmt"

func main() {
	i := 9
	j := 1
	fmt.Println(i, j)
	Swap(&i, &j)
	fmt.Println(i, j)
}

func Swap(i, j *int) {
	*i, *j = *j, *i
}
