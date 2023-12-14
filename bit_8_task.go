package main

import (
	"errors"
	"math"
)

func editBit(num int64, bitnum int64, bitValue bool) (int64, error) {
	err := errors.New("")
	if bitnum > 63 { // проверяем, что номер бита не больше 63
		err = errors.New("bitnum must be less than 64")
		return num, err
	}

	if bitValue { //если включаем бит
		num = num | int64(math.Pow(2, float64(bitnum))) //используем или
	} else { // если выключаем
		if num&int64(math.Pow(2, float64(bitnum))) != 0 { // если бит включен
			num = num ^ int64(math.Pow(2, float64(bitnum))) //используем исключающее или
		}
	}

	return num, err
}

/*func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите число -> ")
	text, _, _ := reader.ReadLine()
	numb, _ := strconv.ParseInt(string(text), 10, 64)

	fmt.Print("Введите номер бита -> ")
	text, _, _ = reader.ReadLine()
	numb2, _ := strconv.ParseInt(string(text), 10, 64)

	fmt.Print("Введите значение бита -> ")
	text, _, _ = reader.ReadLine()
	numb3, _ := strconv.ParseInt(string(text), 10, 64)
	b := false
	if numb3 == 1 {
		b = true
	}

	n, err := editBit(numb, numb2, b)
	fmt.Println(n, err)
}
*/
