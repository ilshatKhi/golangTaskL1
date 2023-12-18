package main

import (
	"sync"
)

// создаем структуру в который есть сам счетчик и mutex для синхронизации
type Counter struct {
	counter int
	mu      sync.Mutex
}

func (c *Counter) increment() {
	//блокируемся
	c.mu.Lock()
	//инкрементируем счетчик
	c.counter++
	//разблокируемся
	c.mu.Unlock()
}

/*func main() {
	//создаем экземпляр Counter-a
	count := Counter{counter: 0}
	//создаем WaitGroup
	var wg sync.WaitGroup
	//запускаем горутины
	for i := 0; i < 100; i++ {
		//увеличиваем счетчик WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.increment()
		}()
	}
	//ожидаем завершения работы горутин
	wg.Wait()

	fmt.Println("Итоговое значение счётчика = ", count.counter)
}*/
