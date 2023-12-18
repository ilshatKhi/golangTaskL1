package main

import (
	"strings"
)

func uniqueChars(input string) bool {
	mapa := make(map[rune]struct{}) //создаем мапу
	//строку делаем в нижнем регистре
	for _, char := range strings.ToLower(input) {
		// если нашли совпадение тогда возвращаем false
		if _, found := mapa[char]; found {
			return false
		}
		mapa[char] = struct{}{}
	}
	return true //если прошли все элементы тогда возвращаем true
}

/*func main() {
	fmt.Println(uniqueChars("abcd"))
	fmt.Println(uniqueChars("abCdefc"))
	fmt.Println(uniqueChars("aabcd"))
}*/
