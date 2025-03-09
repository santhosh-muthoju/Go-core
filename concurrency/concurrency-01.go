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
	time.Sleep(2 * time.Second)
	res := num1 - num2
	fmt.Println("Difference:", res)
	doneChan <- true
	close(doneChan)
}

func PrintInMain01() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go AddNums(40, 20, ch1)
	go SubNums(20, 10, ch2)

	//receiving the value but not pritning ("true")
	<-ch1
	<-ch2

	// If you need true to be printed
	// fmt.Println(<-ch1)
	// fmt.Println(<-ch2)
}
