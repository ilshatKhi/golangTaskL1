package main

import (
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

/*func main() {
	p1 := NewPoint(3, 4)
	p2 := NewPoint(6, 8)

	dist := distance(p1, p2)
	fmt.Printf("Расстояние между точкой p1 и p2: %.2f\n", dist)
}*/
