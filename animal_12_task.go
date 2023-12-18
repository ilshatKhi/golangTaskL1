package main

func setFromarray(sl []string) []string {
	//создаем мапу
	set := make(map[string]string)
	//проходимся по массиву
	for _, val := range sl {
		set[val] = val
	}
	//создали слайс для результата
	var result []string
	//заполнили слайс result
	for k := range set {
		result = append(result, k)
	}
	return result
}

/*func main() {
	fmt.Println(setFromarray([]string{"cat", "cat", "dog", "cat", "tree"}))
}*/
