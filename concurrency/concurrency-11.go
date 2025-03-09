package concurrency

import (
	"fmt"
	"sync"
)

func CounterInc() {
	const numGoRs = 5
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex
	var cond = sync.NewCond(&mu)

	wg.Add(numGoRs)
	currentGoR := 1

	for i := 1; i <= numGoRs; i++ {
		func(id int) {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()

			for id != currentGoR {
				cond.Wait()
			}
			counter++
			fmt.Printf("Goroutine %d incremented counter to %d\n", id, counter)

			currentGoR++
			cond.Broadcast()
		}(i)
	}
	wg.Wait()
	fmt.Println("All goroutines completed.")
}
