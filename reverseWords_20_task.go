package main

import (
	"strings"
)

func reverseWords(input string) string {
	podstr := strings.Fields(input) // разбиваем строку на подстроки по пробелу
	result := ""
	for i := len(podstr) - 1; i >= 0; i-- { // идем в обратную сторону по массиву подстрок
		result += podstr[i] + " " // собираем все в одну строку
	}
	return result
}

/*func main() {
	input := "snow dog sun"
	output := reverseWords(input)
	fmt.Println(output) // Вывод: "sun dog snow"
}
*/
