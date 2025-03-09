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

func PrintMain02() {
	var wg sync.WaitGroup
	ch := make(chan bool)

	wg.Add(3)
	go PrintNums2(ch, &wg)
	go PrintNumsSlow(ch, &wg)
	go PrintNums2(ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for range ch {
	}

	fmt.Println("All goroutines have finished.")
}
