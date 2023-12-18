package main

func intersection(sl1 []int, sl2 []int) []int {
	//алгоритм такой: создаем мапу добавляем туда все, что есть в первом множестве ключ это элемент множества,
	//а значение сколько вхождений данного элемента в мапу
	m := make(map[int]int)
	result := []int{}
	for _, val := range sl1 {
		m[val] = 1
	}
	//здесь добавляем все что есть во втором множестве, если есть пересечение увеличиваем значение мапы на 1
	for _, val := range sl2 {
		if mapVal, ok := m[val]; ok {
			m[val] = mapVal + 1
		}
	}
	// проходимся по мапе и смотрим, где значение == двум записываем в слайс result
	for key, val := range m {
		if val == 2 {
			result = append(result, key)
		}
	}

	return result
}

/*func main() {
	fmt.Println(intersection([]int{1, 2, 3, 4, 6}, []int{0, 2, 5, 6, 1}))
}
*/
