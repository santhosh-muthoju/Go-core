package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func PrintNumsSlow(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 4; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println(i)
		ch <- true
	}
}

func PrintNums2(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 4; i < 7; i++ {
		fmt.Println(i)
		ch <- true
	}
}
