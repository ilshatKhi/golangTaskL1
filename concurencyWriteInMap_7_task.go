package main

import (
	"fmt"
	"sync"
)

func writeInMap(ch <-chan int, mapa map[int]int, mu *sync.RWMutex, wg *sync.WaitGroup) { // здесь принимаем данные по каналу
	for {
		select {
		case value, ok := <-ch: // проверяем не закрыт ли канал
			if !ok { //если закрыт
				fmt.Println("goroutine done")
				wg.Done() //уменьшаем счетчик, чтобы не ждать вечно
				return
			}
			mu.Lock()           // блокируем чтобы записать
			mapa[value] = value // пишем
			mu.Unlock()         // разблокируемся чтоб другие горутины могли писать
		}
	}
}

/*func main() {
	wg := &sync.WaitGroup{}   // создаем waitgroup для ожидания завершения работы всех горутин
	ch := make(chan int, 5)   // канал для передачи данных в горутину
	m := make(map[int]int)    // мапа в которую будем писать
	mu := &sync.RWMutex{}     // рвмьютекс для блокировки горутины, рв здесь не обязательно, но если еще будем читать, то понадобится
	for i := 0; i < 20; i++ { // запускаем 20 горутин
		wg.Add(1) // увеличиваем счетчик
		go writeInMap(ch, m, mu, wg)
	}

	for i := 0; i < 100; i++ { // отсюда пишем в канал
		ch <- i
	}
	close(ch)      //закрываем канал
	wg.Wait()      // ожидаем завершения работы горутин
	fmt.Println(m) // выводим содержимое мапы
}*/
