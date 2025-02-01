package concurrency

// A producer produces items and adds them to a shared buffer.
// A consumer consumes items from the buffer.
// The consumer waits if the buffer is empty, and the producer signals the consumer when new items are available.

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type SharedBuffer struct {
// 	items []int
// 	cond  *sync.Cond
// }

// func NewSharedBuffer() *SharedBuffer {
// 	return &SharedBuffer{
// 		items: make([]int, 0),
// 		cond:  sync.NewCond(&sync.Mutex{}), //return a condition with also a Lock()
// 	}
// }

// func (b *SharedBuffer) Produce(item int) {
// 	b.cond.L.Lock()
// 	defer b.cond.L.Unlock()

// 	// Add item to the buffer
// 	b.items = append(b.items, item)
// 	fmt.Printf("Produced: %d\n", item)

// 	// Signal a waiting consumer
// 	b.cond.Signal()
// }

// func (b *SharedBuffer) Consume() int {
// 	b.cond.L.Lock()
// 	defer b.cond.L.Unlock()

// 	// Wait until there is an item in the buffer
// 	for len(b.items) == 0 {
// 		fmt.Println("Buffer is empty. Consumer is waiting...")
// 		b.cond.Wait() // Releases the lock and waits for a signal
// 	}

// 	// Consume the first item from the buffer
// 	item := b.items[0]
// 	b.items = b.items[1:]
// 	fmt.Printf("Consumed: %d\n", item)

// 	return item
// }

// func PrintInMain08() {
// 	buffer := NewSharedBuffer()

// 	// Producer goroutine
// 	go func() {
// 		for i := 1; i <= 5; i++ {
// 			buffer.Produce(i)
// 			time.Sleep(500 * time.Millisecond) // Simulate production delay
// 		}
// 	}()

// 	// Consumer goroutine
// 	go func() {
// 		for i := 1; i <= 5; i++ {
// 			buffer.Consume()
// 			time.Sleep(1 * time.Second) // Simulate consumption delay
// 		}
// 	}()

// 	// Wait for goroutines to finish
// 	time.Sleep(6 * time.Second)
// 	fmt.Println("Done!")
// }

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type SharedBuffer struct {
	items []int
	cond  *sync.Cond
	ready bool // Indicates whether the buffer is ready for consumption
}

func NewSharedBuffer() *SharedBuffer {
	return &SharedBuffer{
		items: make([]int, 0),
		cond:  sync.NewCond(&sync.Mutex{}),
	}
}

func (b *SharedBuffer) Produce(item int) {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()

	// Wait until the previous item is consumed
	for b.ready {
		fmt.Println("Producer is waiting...")
		b.cond.Wait()
	}

	// Add item to the buffer
	b.items = append(b.items, item)
	fmt.Printf("Produced: %d\n", item)

	// Mark the buffer as ready for consumption
	b.ready = true

	// Signal the consumer
	b.cond.Signal()
}

func (b *SharedBuffer) Consume() int {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()

	// Wait until there is an item in the buffer
	for !b.ready {
		fmt.Println("Consumer is waiting...")
		b.cond.Wait()
	}

	// Consume the first item from the buffer
	item := b.items[0]
	b.items = b.items[1:]
	fmt.Printf("Consumed: %d\n", item)

	// Mark the buffer as empty
	b.ready = false

	// Signal the producer
	b.cond.Signal()

	return item
}

func PrintInMain08() {
	buffer := NewSharedBuffer()

	// Producer goroutine
	go func() {
		for i := 1; i <= 5; i++ {
			buffer.Produce(i)
			time.Sleep(500 * time.Millisecond) // Simulate production delay
		}
	}()

	// Consumer goroutine
	go func() {
		for i := 1; i <= 5; i++ {
			buffer.Consume()
			time.Sleep(1 * time.Second) // Simulate consumption delay
		}
	}()

	// Wait for goroutines to finish
	time.Sleep(10 * time.Second)
	fmt.Println("Done!")
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker canceled:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func PrintInMain09() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go worker(ctx)
	time.Sleep(3 * time.Second)
}
