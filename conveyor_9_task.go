package main

import (
	"fmt"
)

func readAndWrite(ch <-chan int, ch2 chan<- int) {
	for {
		select {
		case value, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(value)
			ch2 <- value * 2
		}
	}
}

/*func main() {
	sl := [100]int{}
	for i := 0; i < 100; i++ {
		sl[i] = i
	}

	ch := make(chan int, 5)
	ch2 := make(chan int, 5)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		readAndWrite(ch, ch2)
	}()

	go func() {
		for val, _ := range sl {
			ch <- val
		}
		close(ch)
	}()

	go func() {
		wg.Wait()
		close(ch2)
	}()

	for value := range ch2 {
		fmt.Println(value)
		_ = value
	}

}*/
