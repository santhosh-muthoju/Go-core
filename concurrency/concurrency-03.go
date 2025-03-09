package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// locking
var mu sync.Mutex

// increasing the counter value to 10 value securely
func IncrementOp(counter *int, wg *sync.WaitGroup) {
	defer wg.Done()

	//for i := 0; i < 10; i++ {
	mu.Lock()
	*counter++
	mu.Unlock()

	time.Sleep(1 * time.Second)
	//}
}

func CallInMain() {
	var wg sync.WaitGroup
	var num int
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go IncrementOp(&num, &wg)
		fmt.Printf("counter value incremented to : %d\n", i)
	}
	wg.Wait()
	fmt.Println("The final counter value:", num)
}
