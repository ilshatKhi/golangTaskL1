package main

func binarySearch(sl []int, reqNum int) int {
	right := len(sl) - 1
	left := 0
	for i := left; i < right; i++ {
		mid := (right + left) / 2
		if sl[mid] == reqNum {
			return mid
		} else {
			if sl[mid] < reqNum {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return 0
}

/*func main() {
	fmt.Println(binarySearch([]int{0,1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 15))
}
*/
