package main

import (
	"sync"
)

type Counter struct {
	counter int
	mu      sync.Mutex
}

func (c *Counter) increment() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

/*func main() {
	count := Counter{counter: 0}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.increment()
		}()
	}

	wg.Wait()

	fmt.Println("Итоговое значение счётчика = ", count.counter)
}*/
