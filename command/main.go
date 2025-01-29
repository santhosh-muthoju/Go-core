package main

import (
	"interviewPrep/concurrency"
)

func main() {
	done := make(chan bool)
	go concurrency.AddNums(10, 20, done)
	go concurrency.SubNums(30, 20, done)
	go concurrency.AddNums(10, 20, done)
	for range done {
	}
}
