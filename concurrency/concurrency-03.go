package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// locking
var mu sync.Mutex

// increasing the counter value to 10 value securley
func IncrementOp(counter *int, ab *sync.WaitGroup) {
	defer ab.Done()

	for i := 0; i < 10; i++ {
		mu.Lock()
		*counter++
		mu.Unlock()

		time.Sleep(1 * time.Millisecond)
	}
}

func CallInMain() {
	var wg sync.WaitGroup
	var num int
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go IncrementOp(&num, &wg)
	}
	wg.Wait()
	fmt.Println("The final counter value:", num)
}
