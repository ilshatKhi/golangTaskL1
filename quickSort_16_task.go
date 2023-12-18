package main

func quickSort(arr []int) []int {
	//определили базовый случай
	if len(arr) <= 1 {
		return arr
	}
	// выбрали опорный элемент посередине масиива
	pivot := arr[(len(arr)-1)/2]
	//объявили слайсы для элементов left которые < pivot и right которые > pivot
	var left, right, pivotSl []int
	//в цикле разделяем на два подмассива
	for i := 0; i < len(arr); i++ {
		if arr[i] < pivot {
			//если элемент < pivot кладем его в левый массив
			left = append(left, arr[i])
		} else if arr[i] > pivot {
			//если элемент > pivot кладем его в левый массив
			right = append(right, arr[i])
		} else if arr[i] == pivot {
			//если элемент == pivot кладем его в левый массив
			pivotSl = append(pivotSl, arr[i])
		}
	}
	//повторяем до наступления базового случая
	left = quickSort(left)
	right = quickSort(right)
	//возвращаем отсортированный массив
	return append(append(left, pivotSl...), right...)
}

/*func main() {
	arr := []int{5, 3, 8, 13, 12, 6, 2, 7, 1, 4, 9, 11, 10, 15, 14}
	sortedArr := quickSort(arr)
	fmt.Println(sortedArr)
}*/
