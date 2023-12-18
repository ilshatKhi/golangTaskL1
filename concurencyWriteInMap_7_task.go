package main

import (
	"fmt"
	"sync"
)

// здесь принимаем данные по каналу
func writeInMap(ch <-chan int, mapa map[int]int, mu *sync.Mutex, wg *sync.WaitGroup) {
	for {
		select {
		// проверяем не закрыт ли канал
		case value, ok := <-ch:
			if !ok { //если закрыт
				fmt.Println("goroutine done")
				//уменьшаем счетчик
				wg.Done()
				return
			}
			// блокируем чтобы записать
			mu.Lock()
			// пишем
			mapa[value] = value
			// разблокируемся чтоб другие горутины могли писать
			mu.Unlock()
		}
	}
}

/*func main() {
	// создаем waitgroup для ожидания завершения работы всех горутин
	wg := &sync.WaitGroup{}
	// канал для передачи данных в горутину
	ch := make(chan int, 5)
	// мапа в которую будем писать
	m := make(map[int]int)
	// рвмьютекс для блокировки горутины, рв здесь не обязательно, но если еще будем читать, то понадобится
	mu := &sync.Mutex{}
	// запускаем 20 горутин
	for i := 0; i < 20; i++ {
		// увеличиваем счетчик
		wg.Add(1)
		go writeInMap(ch, m, mu, wg)
	}

	// отсюда пишем в канал
	for i := 0; i < 100; i++ {
		ch <- i
	}
	//закрываем канал
	close(ch)
	// ожидаем завершения работы горутин
	wg.Wait()
	// выводим содержимое мапы
	fmt.Println(m)
}*/
