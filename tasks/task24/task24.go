package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде
// структуры Point с инкапсулированными параметрами x, y и конструктором.

type Point struct {
	x, y int
}

func MakePoint(x, y int) Point {
	return Point{x: x, y: y}
}

func GetDist(a, b Point) float64 {
	return math.Sqrt(float64(
		(a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y),
	))
}

func main() {
	a := MakePoint(39, -12)
	b := MakePoint(-54, 17)

	fmt.Println(GetDist(a, b))
}
