package main

import "fmt"

func typeAssertion(t interface{}) {
	//проверяем на chan
	if v, ok := t.(chan int); ok {
		fmt.Println(v, "is chan int")
	}
	//проверяем на int
	if v, ok := t.(int); ok {
		fmt.Println(v, "is int64")
	}
	//проверяем на bool
	if v, ok := t.(bool); ok {
		fmt.Println(v, "is bool")
	}
	//проверяем на string
	if v, ok := t.(string); ok {
		fmt.Println(v, "is string")
	}
}

/*func main() {
	ch := make(chan int)
	typeAssertion(ch)
	typeAssertion(64)
	typeAssertion(true)
	typeAssertion("Это строка")
}*/
