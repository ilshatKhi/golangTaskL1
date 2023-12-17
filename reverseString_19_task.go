package main

func reverseString(input string) string {
	runes := []rune(input)                                // конвертим в rune чтобы можно было выводить символы unicode
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { //
		runes[i], runes[j] = runes[j], runes[i] // тут меняем местами символы первый и последний и движемся к центру
	}
	return string(runes)
}

/*func main() {
	inputString := "абвгдейка"
	reversedString := reverseString(inputString)
	fmt.Println(reversedString)
}*/
