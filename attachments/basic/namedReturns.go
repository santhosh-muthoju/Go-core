package basic

import "fmt"

// Named return values: sum and diff
func calculate(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return // no need to specify variables explicitly
}

func PrintNamedReturns() {
	s, d := calculate(10, 5)
	fmt.Println("Sum:", s)
	fmt.Println("Difference:", d)
}
