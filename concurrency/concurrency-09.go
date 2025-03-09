package concurrency

import (
	"fmt"
	"time"
)

func NumOps() {
	num1 := 60
	num2 := 50

	AddNumbers(num1, num2)
	AddNumbers(num1, num2)

	resultChan := make(chan int)
	go SubNumbers(num1, num2, resultChan)

	AddNumbers(num1, num2)

	result := <-resultChan
	fmt.Println(result)
}

func AddNumbers(num1, num2 int) {
	res := num1 + num2
	fmt.Println(res)
}

func SubNumbers(num1, num2 int, done chan int) {
	time.Sleep(2 * time.Second)
	res := num1 - num2
	done <- res
}
