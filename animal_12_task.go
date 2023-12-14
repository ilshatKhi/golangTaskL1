package main

func setFromarray(sl []string) []string {
	m := make(map[string]string)
	for _, val := range sl {
		m[val] = val
	}
	var newsl []string
	for k := range m {
		newsl = append(newsl, k)
	}
	return newsl
}

/*func main() {
	fmt.Println(setFromarray([]string{"cat", "cat", "dog", "cat", "tree"}))
}*/
