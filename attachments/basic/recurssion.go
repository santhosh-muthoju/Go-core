package basic

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func PrintInMainFact() {
	fmt.Println(fact(7))

	var fib func(n int) int
	fib = func(n int) int {
		if n == 0 {
			return 0 // fib(0) should return 0
		}
		if n == 1 {
			return 1 // fib(1) should return 1
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}
