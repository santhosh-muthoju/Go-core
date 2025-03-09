package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func PrintInMain10() {
	fmt.Println("Main function is running...")
	var wg sync.WaitGroup

	wg.Add(1)
	go PrintNumbers(&wg) // Start a goroutine without a channel

	// Wait for the goroutine to finish (for demonstration purposes)
	//time.Sleep(6 * time.Second)
	wg.Wait()
	fmt.Println("Main function exiting.")
}

func PrintNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
