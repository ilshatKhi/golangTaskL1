package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ctxWithDeadLine() {
	//создаем канал для записи данных
	ch := make(chan int, 5)
	//создаем контекст с дедлайном
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	defer cancel()
	//создаем waitgroup
	wg := &sync.WaitGroup{}
	//увеличили счетчик waitgroup
	wg.Add(1)
	//запускаем work
	go work(ctx, ch, wg)
	//запускаем пишущую горутину
	go func() {
		defer close(ch)
		i := 0
		for {
			select {
			case <-ctx.Done(): // отслеживаем контекст
				fmt.Println("<-ctx.Done():")
				return // выходим из горутины
			default:
				//пишем в канал
				ch <- i
				i++
				//моделируем полезную работу
				time.Sleep(time.Millisecond * 1)
			}
		}
	}()
	wg.Wait() // ждем завершения выполнения горутин
}

func work(ctx context.Context, inputChanel <-chan int, group *sync.WaitGroup) {
	for {
		select {
		// если пришел Done - выходим
		case <-ctx.Done():
			//уменьшаем счетчик
			group.Done()
			fmt.Println("ctx.Done()")
			return
		default:
			//читаем и выводим то, что пришло в канал
			fmt.Println(<-inputChanel)
		}

	}
}

/*func main() {
	ctxWithDeadLine()
}*/
