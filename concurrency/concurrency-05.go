package concurrency

import (
	"fmt"
	"time"
)

func PrintInMain05() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Hello from channel 1"
	}()

	// as this ch2 with less time i.e 1 sec will execute faster and gives it as o/p
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Hello from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout! No message received")
	}

}
