package main

import "fmt"

var justString string

func createHugeString(n int) string {
	return "asdkalksnflansdlfknaslkdnflaksndflaksndflkasndflknasldfknaslkjjkljkljkljkljkljkljkljkljkljklljkkkdfn"
}

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) >= 100 {
		justString = v[:100]
	} else {
		fmt.Println("v is short")
		justString = ""
	}

	fmt.Println(justString)
}

/*func main() {
	someFunc()
}*/
