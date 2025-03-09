package concurrency

import (
	"fmt"
	"sync"
)

func CounterInc2() {
	const numGoroutines = 5 // Number of goroutines
	var counter int
	var wg sync.WaitGroup
	ch := make(chan int, 1) // Buffered channel with size 1 to control execution order

	wg.Add(numGoroutines)

	// Start with goroutine 1
	ch <- 1

	for i := 1; i <= numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			<-ch // Wait for turn

			// Increment counter and print
			counter++
			fmt.Printf("Goroutine %d incremented counter to %d\n", id, counter)

			// Pass control to next goroutine
			if id < numGoroutines {
				ch <- id + 1
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines completed.")
}
