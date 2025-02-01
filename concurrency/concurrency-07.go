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

func PrintINMain07(num *Singleton) {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		func() {
			defer wg.Done()
			num = GetSingleton()
			fmt.Printf("Singleton value: %d\n", num.num)
		}()
	}
	wg.Wait()

	// var dummy *concurrency.Singleton
	// concurrency.PrintINMain07(dummy)
}
