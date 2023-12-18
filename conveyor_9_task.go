package main

/*func main() {
	// Создаем каналы для передачи данных
	inputChan := make(chan int)
	outputChan := make(chan int)
	//создаем массив
	numbers := [100]int{}
	//заполняем массив
	for i := 0; i < 100; i++ {
		numbers[i] = i
	}

	// Горутина для записи чисел в первый канал
	go func() {
		for _, num := range numbers {
			inputChan <- num
		}
		//закрываем канал после записи
		close(inputChan)
	}()

	// Горутина для умножения чисел на 2 и записи в outputChan
	go func() {
		for num := range inputChan {
			outputChan <- num * 2
		}
		//закрываем канал
		close(outputChan)
	}()

	// Вывод результата
	for result := range outputChan {
		fmt.Println(result)
	}
}*/
