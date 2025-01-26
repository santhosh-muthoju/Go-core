package main

import (
	"fmt"
	"time"
)

func numberPrint(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}

func main() {
	go numberPrint("1") // new go routine
	numberPrint("2")    // main go routine
}
