package main

import "fmt"

type Human struct {
	Name string
	Age  int
	test string
}

type Action struct {
	Human
	test string
}

func (h *Human) Hi() {
	fmt.Println("Hello from Human")
}

func (h *Human) Bye() {
	fmt.Printf("Bye from %s\n", h.Name)
}

func (a *Action) Hi() {
	fmt.Println("Hello from Action")
}

func main() {
	a := Action{Human: Human{Name: "Dima", Age: 22, test: "test-dima"}}
	b := Action{Human: Human{Name: "Vadim", Age: 15, test: "test-vadim"}, test: "test-updated"}
	fmt.Println(a.Name, b.Name)
	fmt.Println(a.Age, b.Age)
	fmt.Println(a.test, b.test)
	a.Hi()
	b.Hi()
	a.Human.Bye()
	b.Bye()
}
