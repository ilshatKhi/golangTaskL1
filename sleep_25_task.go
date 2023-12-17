package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	select {
	case <-time.After(time.Second * 5):
		fmt.Println("timed out")
	}
}

/*func main() {
	fmt.Println("До sleep")
	sleep(time.Second * 5)
	fmt.Println("После sleep")
}*/
