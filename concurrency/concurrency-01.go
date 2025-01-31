package concurrency

import (
	"fmt"
	"time"
)

func AddNums(num1, num2 int, doneChan chan bool) {
	res := num1 + num2
	fmt.Println("sum:", res)
	doneChan <- true
}

func SubNums(num1, num2 int, doneChan chan bool) {
	time.Sleep(1 * time.Second)
	res := num1 - num2
	fmt.Println("Difference:", res)
	doneChan <- true
	close(doneChan)
}
