package main

import (
	"strings"
)

func reverseWords(input string) string {
	// разбиваем строку на подстроки по пробелу
	strFields := strings.Fields(input)
	result := strings.Builder{}
	// идем в обратную сторону по массиву подстрок
	for i := len(strFields) - 1; i >= 0; i-- {
		result.WriteString(strFields[i])
		result.WriteString(" ")
	}
	return result.String()
}

/*func main() {
	input := "snow dog sun"
	output := reverseWords(input)
	fmt.Println(output)
}
*/
