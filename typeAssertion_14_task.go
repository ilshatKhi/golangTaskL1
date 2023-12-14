package main

import "fmt"

func typeAssertion(t interface{}) {
	if v, ok := t.(int); ok {
		fmt.Println(v, "is int64")
	}

	if v, ok := t.(bool); ok {
		fmt.Println(v, "is bool")
	}

	if v, ok := t.(string); ok {
		fmt.Println(v, "is string")
	}
}

/*func main() {
	typeAssertion(64)
	typeAssertion(true)
	typeAssertion("kajshndjkla")
}
*/
