package main

import (
	"errors"
	"math"
)

func editBit(num int64, bitnum int64, bitValue bool) (int64, error) {
	err := errors.New("")
	// проверяем, что номер бита не больше 63
	if bitnum > 63 {
		err = errors.New("bitnum must be less than 64")
		return num, err
	}

	if bitValue { //если включаем бит
		//используем или
		num = num | int64(math.Pow(2, float64(bitnum)))
	} else {
		if num&int64(math.Pow(2, float64(bitnum))) != 0 {
			//используем xor
			num = num ^ int64(math.Pow(2, float64(bitnum)))
		}
	}

	return num, err
}

/*func main() {
	reader := bufio.NewReader(os.Stdin)
	//вводим число в котором меняем бит
	fmt.Print("Введите число -> ")
	text, _, _ := reader.ReadLine()
	numb, _ := strconv.ParseInt(string(text), 10, 64)
	//вводим номер бита от 0 до 63 включительно
	fmt.Print("Введите номер бита 0<=number<64 -> ")
	text, _, _ = reader.ReadLine()
	numb2, _ := strconv.ParseInt(string(text), 10, 64)
	if numb2 > 63 || numb2 < 0 {
		fmt.Println("bitnum must be less than 64")
		return
	}
	//вводим значение бита
	fmt.Print("Введите значение бита -> ")
	text, _, _ = reader.ReadLine()
	numb3, _ := strconv.ParseInt(string(text), 10, 64)
	b := false
	if numb3 >= 1 {
		b = true
	}
	//
	n, err := editBit(numb, numb2, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Число после смены бита - ", n)
}*/
