package concurrency

import (
	"fmt"
	"time"
)

func TestWorker(id int, errCh chan error) {
	time.Sleep(2 * time.Second)

	if id%2 == 0 {
		errCh <- fmt.Errorf("worker %d encountered an error", id)
	} else {
		errCh <- nil
	}
}
func TestWorkerNew(id int, errCh chan error) {
	time.Sleep(2 * time.Second)

	if id%2 == 0 {
		errCh <- fmt.Errorf("worker %d encountered an error", id)
	} else {
		errCh <- nil
	}
}

func PrintInMain06() {
	errCh := make(chan error, 3)

	for i := 1; i <= 3; i++ {
		go TestWorker(i, errCh)
	}

	for i := 1; i <= 3; i++ {
		select {
		case err := <-errCh:
			if err != nil {
				fmt.Println("Error received", err)
			} else {
				fmt.Println("worker completed succesfully!")
			}
		case <-time.After(10 * time.Second):
			fmt.Println("Timeout! No response received!")
		}
	}

	// for i := 1; i <= 3; i++ {
	// 	err := <-errCh
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 	} else {
	// 		fmt.Println("Worker completed successfully!")
	// 	}
	// }
}
