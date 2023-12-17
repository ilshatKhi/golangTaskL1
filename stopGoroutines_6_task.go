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

/*func main() {
	ch := make(chan bool)   // создали небуферизированный канал для передачи сигнала
	wg := &sync.WaitGroup{} //создаем WaitGroup для синхронизации
	wg.Add(1)
	go secondMethode(ch, wg)
	time.Sleep(time.Second * 3)
	ch <- true

	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go firstMethode(ctx, wg)
	time.Sleep(time.Second * 3)
	cancel()

	wg.Wait()
}*/

func someAction(v []int8, b int8) {
	fmt.Println(v)
	v[0] = 100
	fmt.Println(v)
	v = append(v, b)
	fmt.Println(v)
}

/*func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}*/
