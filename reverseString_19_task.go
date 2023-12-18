package main

func reverseString(input string) string {
	// конвертим в rune чтобы можно было выводить символы unicode
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// тут меняем местами символы первый и последний и движемся к центру
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

/*func main() {
	inputString := "абвгдейка"
	reversedString := reverseString(inputString)
	fmt.Println(reversedString)
}*/
