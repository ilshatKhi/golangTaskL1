package main

import (
	"fmt"
	"sync"
)

func sumConc(sl []int) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	sum := 0
	//fmt.Println(sl)
	for _, num := range sl {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			mu.Lock()
			sum += num * num
			mu.Unlock()
		}(num)
	}
	wg.Wait()
	fmt.Println(sum)
}

/*
func main() {
	sumConc([]int{2, 4, 6, 8, 10})
}*/
