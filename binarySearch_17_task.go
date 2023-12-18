package main

func binarySearch(sl []int, reqNum int) int {
	//запоминаем индекс крайнего правого элемента
	right := len(sl) - 1
	//запоминаем индекс левого элемента
	left := 0
	//запускаем цикл
	//при каждой итерации уменьшаем размер массива в 2 раза
	for i := left; i < right; i++ {
		//берем индекс центрального элемента
		mid := (right + left) / 2
		//если мы нашли искомый элемент, то возвращаем индекс элемента
		if sl[mid] == reqNum {
			return mid
		} else {
			//если он меньше искомого элемента сдвигаем левую границу, иначе правую
			if sl[mid] < reqNum {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	//если не нашли возвращаем -1
	return -1
}

/*func main() {
	mas := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	index := binarySearch(mas, 15)
	if index == -1 {
		fmt.Println("Элемент не найден")
	} else {
		fmt.Println("Индекс искомого элемента =", index)
	}
}*/
