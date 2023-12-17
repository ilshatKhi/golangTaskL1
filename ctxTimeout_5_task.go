package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ctxWithDeadLine() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 5)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		work(ctx, ch)
	}()
	go func() {
		defer close(ch)
		i := 0
	loop:
		for {
			select {
			case <-ctx.Done():
				fmt.Println("case <-ctx.Done():")
				break loop
			default:
				ch <- i
				time.Sleep(time.Millisecond * 1)
			}
		}

	}()
	wg.Wait()
}

func work(ctx context.Context, inputChanel <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx.Done()")
			return
		default:

		}
		fmt.Println(<-inputChanel)
	}
}

/*func main() {
	ctxWithDeadLine()
}*/
