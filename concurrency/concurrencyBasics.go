package concurrency

import (
	"fmt"
	"time"
)

func Greet(str string, doneChan chan bool) {
	fmt.Println("Hello!", str)
	doneChan <- true
}

func SlowGreet(str string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hi!", str)
	doneChan <- true
	close(doneChan)
}

// printing the data received from multiple channels through select satement
func PrintINMain() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "World"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
