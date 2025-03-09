package concurrency

import (
	"fmt"
	"sync"
)

type Singleton struct {
	num int
}

var (
	instance *Singleton
	once     sync.Once
)

func GetSingleton() *Singleton {
	once.Do(func() {
		fmt.Println("Creating single instance:")
		instance = &Singleton{num: 42}
	})
	return instance
}

func PrintINMain07(value *Singleton) {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		func() {
			defer wg.Done()
			value = GetSingleton()
			fmt.Printf("Singleton value: %d\n", value.num)
		}()
	}
	wg.Wait()

	// var dummy *concurrency.Singleton
	// concurrency.PrintINMain07(dummy)
}
