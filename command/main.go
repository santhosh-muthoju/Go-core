package main

import "interviewPrep/concurrency"

func main() {
	done := make(chan bool)
	go concurrency.Greet("Nice to meet you..", done)
	go concurrency.SlowGreet("Watsupppp?", done)
	go concurrency.Greet("I hope you are doing good!", done)
	<-done
	<-done
	<-done
}
