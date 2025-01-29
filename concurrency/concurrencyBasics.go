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
}
