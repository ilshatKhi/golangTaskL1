package main

import "fmt"

var justString string

func createHugeString(n int) string {
	return "asdkalk2snasdkalk2snasdkalk2snasdkalk2snasdkalk2snasdkalk2snasdkalk2snasdkalk2snasdkalk2snasdkalk2snooo"
}

func someFunc() {
	// тут добавил проверку на длину, если длины не хватит то получаем out of range
	v := createHugeString(1 << 10)
	//длину проверяем
	if len([]rune(v)) >= 100 {
		justString = v[:100]
	} else {
		fmt.Println("string is short")
		justString = ""
	}

	fmt.Println(justString)
}

/*func main() {
	someFunc()
}*/
