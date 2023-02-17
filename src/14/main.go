package main

import "fmt"

func main() {
	var v any

	//v = 12
	//v = "zxc"
	//v = make(chan any)
	//v = true
	typeDeterminant(v)
}

func typeDeterminant(v any) {
	switch v.(type) {
	case int:
		fmt.Println("integer type")
	case string:
		fmt.Println("string type")
	case bool:
		fmt.Println("bool type")
	case chan any:
		fmt.Println("channel type")
	}
}
