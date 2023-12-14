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
)

func worker(ctx context.Context, data <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx.Done() закрываем воркер")
			return
		case value, ok := <-data:
			if !ok {
				return
			}
			fmt.Println(value)
		}
	}
}

func gracefulShutdown() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	ch := make(chan int, 10)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите число воркеров: ")
	text2, _, _ := reader.ReadLine()
	workersNum, _ := strconv.ParseInt(string(text2), 10, 64)
	fmt.Printf("Workers count - %d\n", workersNum)

	for i := 0; i < int(workersNum); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, ch)
		}()
	}

	go func(sigch chan os.Signal) {
		i := 0
		for {
			ch <- i
			i++
			select {
			case <-sigChan:
				cancel()
				break
			default:
				continue
			}
		}
	}(sigChan)

	wg.Wait()
}

/*func main() {
	gracefulShutdown()
}*/
