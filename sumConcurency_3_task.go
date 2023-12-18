package main

import (
	"fmt"
	"sync"
)

func sumConc(sl []int) {
	var wg sync.WaitGroup // создаем waitgroup
	var mu sync.Mutex     // создаем mutex, чтоб блокироваться при записи
	sum := 0
	for _, num := range sl { // в цикле запускаем горутины
		wg.Add(1) // увеличиваем счетчик waitgroup
		go func(num int) {
			defer wg.Done()  //уменьшаем
			mu.Lock()        // блокируемся, для того чтобы только 1 горутина одномоментно имела доступ к переменной sum
			sum += num * num // суммируем
			mu.Unlock()      //разблокируемся
		}(num) // передаем число в параметрах
	}
	wg.Wait() //ожидаем завершения работы горутин
	fmt.Println(sum)
}

/*func main() {
	sumConc([]int{2, 4, 6, 8, 10})
}*/
