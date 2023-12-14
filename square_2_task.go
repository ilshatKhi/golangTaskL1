package main

import "fmt"
import "math"

func printSquare(mas []int) {
	for _, num := range mas {
		fmt.Println(math.Pow(float64(num), 2))
	}
}
