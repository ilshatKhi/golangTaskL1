package main

import (
	"math/big"
)

/*
здесь использован пакет big  чтобы можно было умножать числа больше чем 2^20
*/
func multiply(a string, b string) *big.Int {
	result := new(big.Int) //
	first := new(big.Int)
	second := new(big.Int)

	first.SetString(a, 10)
	second.SetString(b, 10)

	result.Mul(first, second)
	return result
}

func division(a string, b string) *big.Int {
	result := new(big.Int)
	first := new(big.Int)
	second := new(big.Int)

	first.SetString(a, 10)
	second.SetString(b, 10)

	result.Div(first, second)
	return result
}

func addition(a string, b string) *big.Int {
	result := new(big.Int)
	first := new(big.Int)
	second := new(big.Int)

	first.SetString(a, 10)
	second.SetString(b, 10)

	result.Add(first, second)
	return result
}

func subtraction(a string, b string) *big.Int {
	result := new(big.Int)
	first := new(big.Int)
	second := new(big.Int)

	first.SetString(a, 10)
	second.SetString(b, 10)

	result.Sub(first, second)
	return result
}

/*func main() {
	a := "10485700000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007"
	b := "2097000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000154"

	fmt.Println("Умножение:", multiply(a, b))
	fmt.Println("Деление:", division(a, b))
	fmt.Println("Сложение:", addition(a, b))
	fmt.Println("Вычитание:", subtraction(a, b))
}*/
