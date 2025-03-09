package concurrency

//basics of RWMutex function // use when many read op's and few write op's
//A write lock blocks all other read and write locks.
//A read lock blocks only write locks.

import (
	"fmt"
	"sync"
	"time"
)

var Data = make(map[string]string)
var RwLock sync.RWMutex

func ReadData(key string, wg *sync.WaitGroup) {
	defer wg.Done()

	RwLock.RLock()         // Acquire a read lock
	defer RwLock.RUnlock() //Release the read lock when done

	value, exists := Data[key]
	if exists {
		fmt.Printf("Read: key=%s, value=%s\n", key, value)
	} else {
		fmt.Printf("Read: key=%s not found\n", key)
	}
	time.Sleep(1 * time.Second)
}

func WriteData(key string, value string, wg *sync.WaitGroup) {
	defer wg.Done()

	RwLock.Lock()         // write lock
	defer RwLock.Unlock() // write unlock

	Data[key] = value
	fmt.Printf("Write: key=%s, value=%s\n", key, value)
}

func CallinMain04() {
	var wg sync.WaitGroup

	wg.Add(1)
	go WriteData("Christiano", "Ronaldo", &wg)

	//now open multiple readers
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go ReadData("Christiano", &wg)
	}

	wg.Add(1)
	go WriteData("Leonel", "Messi", &wg)

	// start more readers
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go ReadData("Christiano", &wg)
	}

	wg.Wait()
	fmt.Println("Sample data:", Data)
}
