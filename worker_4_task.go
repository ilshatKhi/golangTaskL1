package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// создаем функцию воркер
func worker(ctx context.Context, data <-chan int) { //
	for {
		//смотрим если пришел ctx.Done - закрываем воркер, он пришел значит закрыли контекст
		select {
		case <-ctx.Done():
			fmt.Println("ctx.Done() закрываем воркер")
			return
			//читаем из канала и выводим, также проверяем что канал не закрыт
		case value, ok := <-data:
			if !ok {
				return
			}
			fmt.Println(value)
		}
	}
}

func Shutdown() {
	// создаем канал для записи данных
	ch := make(chan int, 10)
	// создаем контекст для завершения работы воркеров
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//создаем waitgroup
	wg := &sync.WaitGroup{}

	// для удобства вводим через консоль кол-во воркеров
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите число воркеров: ")
	text2, _, _ := reader.ReadLine()
	workersNum, _ := strconv.ParseInt(string(text2), 10, 64)
	fmt.Printf("Число воркеров - %d\n", workersNum)
	// чтоб успеть прочитать
	time.Sleep(time.Second * 2)

	// запускаем воркеры
	for i := 0; i < int(workersNum); i++ {
		// увеличиваем счетчик
		wg.Add(1)
		go func() {
			// уменьшаем счетчик
			defer wg.Done()
			// передаем в воркер контекст и канал для данных
			worker(ctx, ch)
		}()
	}

	// запускаем главный поток который записывает данные в горутину
	go func() {
		i := 0
		// запускаем бесконечный цикл
		for {
			// инкрементируем i
			i++
			select {
			// отлавливаем сигнал на завершение
			case <-ctx.Done():
				// отменяем контекст, останавливаем воркеры
				cancel()
				break
			default:
				// отправляем в канал данные
				ch <- i
			}
		}
	}()

	// ждем завершения всех горутин
	wg.Wait()
}

/*func main() {
	Shutdown()
}*/
