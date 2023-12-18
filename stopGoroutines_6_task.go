package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func firstMethode(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done(): //когда приходит пустая структура из канала значит контекст отменен
			wg.Done()
			fmt.Println("goroutine stopped firstMethode")
			return
		default: // если сигнала нет то идем в ветку default
			fmt.Println("goroutine work firstMethode")
			time.Sleep(time.Millisecond * 1)
		}
	}
}

func secondMethode(ch chan bool, wg *sync.WaitGroup) {
	for { // запускаем бесконечный цикл
		select {
		case <-ch: // это блокирующая операция
			wg.Done()
			fmt.Println("goroutine stopped secondMethode") // пришел сигнал на остановку горутины
			return
		default: // если сигнала нет то идем в ветку default
			time.Sleep(time.Millisecond * 1)
			fmt.Println("goroutine work secondMethode")
		}
	}
}

func thirdMethode(wg *sync.WaitGroup, mutex sync.RWMutex) {
	//уменьшаем счетчик
	defer wg.Done()
	//запускаем бесконечный цикл
	for {
		//будем делать RLock потому что мы здесь читаем значение stopped, если нужно менять нужен Lock
		mutex.RLock()
		//флаг остановки тру тогда завершаем горутину
		if stopped {
			//прочли значение stopped делаем  RUnlock
			mutex.RUnlock()
			// выходим из горутины
			fmt.Println("goroutine stopped thirdMethode")
			return
		}
		//прочли значение stopped делаем  RUnlock
		mutex.RUnlock()
		fmt.Println("goroutine work thirdMethode")
		time.Sleep(time.Millisecond * 1)
	}
}

func stopGoroutine3Methode() {
	stopped = true
}

var stopped bool

/*func main() {
	//создаем WaitGroup для синхронизации
	wg := &sync.WaitGroup{}
	//в первом способе использован контекст
	//создаем контекст который отменим
	ctx, cancel := context.WithCancel(context.Background())
	//увеличиваем счетчик
	wg.Add(1)
	//запускам горутину
	go firstMethode(ctx, wg)
	//спим секунду
	time.Sleep(time.Second * 1)
	//отменяем контекст этим выполняем останов горутины
	cancel()

	//во втором способе использован канал для передачи сигнала останова
	//создали канал для передачи сигнала останова
	ch := make(chan bool)
	//увеличиваем счетчик
	wg.Add(1)
	//запускаем горутину
	go secondMethode(ch, wg)
	//спим секунду
	time.Sleep(time.Second * 1)
	//отправляем true в канал и останавливаем горутину
	ch <- true

	//третий способ использует mutex и булевый глобальный флаг
	//создаем mutex для блокировки при чтении
	mutex := sync.RWMutex{}
	//увеличиваем счетчик
	wg.Add(1)
	//запускаем горутину
	go thirdMethode(wg, mutex)
	//спим секунду
	time.Sleep(time.Second * 1)
	// устанавливаем флаг остановки
	stopGoroutine3Methode()

	//ждем завершения работы горутин
	wg.Wait()
}*/
