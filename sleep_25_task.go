package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	//реализовано с помощью time.After(time.Second * 5) которая возвращает
	select {
	//пока не придет канал мы блокируемся здесь
	case <-time.After(time.Second * 5):
		fmt.Println("timed out")
	}
}

func main() {
	fmt.Println("До sleep")
	sleep(time.Second * 5)
	fmt.Println("После sleep")
}
