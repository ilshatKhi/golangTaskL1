package main

import (
	"strings"
)

func uniqueChars(input string) bool {
	mapa := make(map[rune]struct{}) //создаем мапу

	for _, char := range strings.ToLower(input) {
		if _, found := mapa[char]; found { // если нашли совпадение тогда возвращаем тру
			return false
		}
		mapa[char] = struct{}{}
	}
	return true //если прошли все элементы тогда возвращаем тру
}

/*func main() {
	fmt.Println(uniqueChars("abcd"))
	fmt.Println(uniqueChars("abCdefc"))
	fmt.Println(uniqueChars("aabcd"))
}*/
