package concurrency

import (
	"fmt"
	"time"
)

func Greet(str string, doneChan chan string) {
	fmt.Println("Hello!", str)
	doneChan <- "Greeting done!"
}

func SlowGreet(str string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hi!", str)
	doneChan <- true
	close(doneChan)
}

func PrintInMainBasic() {
	str := "santhosh"
	ch := make(chan string)
	ch1 := make(chan bool)

	go Greet(str, ch)
	go SlowGreet(str, ch1)

	greetResponse := <-ch
	fmt.Println(greetResponse)

	slowGreetDone := <-ch1
	fmt.Println("slow greet completed:", slowGreetDone)

}

// printing the data received from multiple channels through select satement
func PrintINMain() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "World"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
