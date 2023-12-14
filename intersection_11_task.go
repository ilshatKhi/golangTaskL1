package main

func intersection(sl1 []int, sl2 []int) []int {
	m := make(map[int]int)
	sl3 := []int{}
	for _, val := range sl1 {
		m[val] = 1
	}

	for _, val := range sl2 {
		if mapVal, ok := m[val]; ok {
			m[val] = mapVal + 1
		}
	}

	for key, val := range m {
		if val == 2 {
			sl3 = append(sl3, key)
		}
	}

	return sl3
}

/*func main() {
	fmt.Println(intersection([]int{1, 2, 3, 4, 6}, []int{0, 2, 5, 6, 1}))
}
*/
