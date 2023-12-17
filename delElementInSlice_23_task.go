package main

func delElement(sl []int, i int) []int {
	var err error
	if i >= 0 && i < len(sl) {
		sl = append(sl[:i], sl[i+1:]...)
	} else {
		err.Error()
	}
	return sl
}

/*func main() {
	fmt.Println(delElement([]int{1, 2, 3, 4, 5, 6, 7}, 9))
}
*/
