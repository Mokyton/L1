package main

import (
	"24/point"
	"fmt"
	"math"
)

func findDistance(a, b *point.Point) float64 { // Формула нахождения расстояния между точками на плоскости
	xA := a.GetX()
	xB := b.GetX()
	yA := a.GetY()
	yB := b.GetY()
	return math.Sqrt(math.Pow(xB-xA, 2) + math.Pow(yB-yA, 2))
}

func main() {
	a := point.NewPoint(-1, 3)
	b := point.NewPoint(6, 2)
	fmt.Println(findDistance(a, b))
	fmt.Println(findDistance(b, a))
	a.SetX(0)
	a.SetY(1)
	b.SetY(-2)
	b.SetX(2)
	fmt.Println(findDistance(a, b))
	fmt.Println(findDistance(b, a))
}
